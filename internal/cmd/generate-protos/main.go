// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run -tags protolegacy . -execute

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"

	gengo "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/internal/detrand"
	"google.golang.org/protobuf/internal/editionssupport"
)

func init() {
	// Determine repository root path.
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").CombinedOutput()
	check(err)
	repoRoot = strings.TrimSpace(string(out))

	// Determine the module path.
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Path}}")
	cmd.Dir = repoRoot
	out, err = cmd.CombinedOutput()
	check(err)
	modulePath = strings.TrimSpace(string(out))

	// When the environment variable RUN_AS_PROTOC_PLUGIN is set,
	// we skip running main and instead act as a protoc plugin.
	// This allows the binary to pass itself to protoc.
	if plugin := os.Getenv("RUN_AS_PROTOC_PLUGIN"); plugin != "" {
		// Disable deliberate output instability for generated files.
		// This is reasonable since we fully control the output.
		detrand.Disable()

		protogen.Options{}.Run(func(gen *protogen.Plugin) error {
			for _, file := range gen.Files {
				if file.Generate {
					gengo.GenerateVersionMarkers = false
					gengo.GenerateFile(gen, file)
					generateIdentifiers(gen, file)
					generateSourceContextStringer(gen, file)
				}
			}
			gen.SupportedFeatures = gengo.SupportedFeatures
			return nil
		})
		os.Exit(0)
	}
}

var (
	run        bool
	protoRoot  string
	repoRoot   string
	modulePath string

	generatedPreamble = []string{
		"// Copyright 2019 The Go Authors. All rights reserved.",
		"// Use of this source code is governed by a BSD-style",
		"// license that can be found in the LICENSE file.",
		"",
		"// Code generated by generate-protos. DO NOT EDIT.",
		"",
	}
)

func main() {
	flag.BoolVar(&run, "execute", false, "Write generated files to destination.")
	flag.StringVar(&protoRoot, "protoroot", os.Getenv("PROTOBUF_ROOT"), "The root of the protobuf source tree.")
	flag.Parse()
	protocPath, err := exec.LookPath("protoc")
	if err != nil {
		panic("protoc not found in $PATH.")

	}
	if !strings.Contains(protocPath, ".cache/bin") {
		fmt.Fprintf(os.Stderr, "found protoc in $PATH but it is not in '.cache/bin`.\nRun ./test.bash and add '.cache/bin' to your $PATH to make sure you are using the same protoc as ./test.bash.\n")
	}
	if protoRoot == "" {
		panic("protobuf source root is not set")
	}

	// Generate editions_defaults.binpb first before generating any code for
	// protos: the .proto files might specify a very recent edition for which
	// editions_default.binpb was not yet updated.
	generateEditionsDefaults()

	// Generate versions of each testproto .proto file which use the Hybrid and
	// Opaque API. This step needs to come first so that the next step will
	// generate the .pb.go files for these extra .proto files.
	generateOpaqueTestprotos()

	generateLocalProtos()
	generateRemoteProtos()
}

// gsed works roughly like sed(1), in that it processes a file with a list of
// replacement functions that are applied in order to each line.
func gsed(outFn, inFn string, repls ...func(line string) string) error {
	if err := os.MkdirAll(filepath.Dir(outFn), 0755); err != nil {
		return err
	}
	out, err := os.Create(outFn)
	if err != nil {
		return err
	}
	defer out.Close()
	b, err := os.ReadFile(inFn)
	if err != nil {
		return err
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	for idx, line := range lines {
		for _, repl := range repls {
			line = repl(line)
		}
		lines[idx] = line
	}
	if _, err := out.Write([]byte(strings.Join(lines, "\n"))); err != nil {
		return err
	}
	return out.Close()
}

// variantFn turns a relative path like
// internal/testprotos/annotation/annotation.proto into its corresponding
// Hybrid/Opaque API variant file name,
// e.g. internal/testprotos/annotation/annotation_hybrid/annotation.hybrid.proto
func variantFn(relPath, variant string) string {
	base := strings.TrimSuffix(filepath.Base(relPath), ".proto")
	dir := filepath.Dir(relPath)
	return filepath.Join(dir, filepath.Base(dir)+"_"+variant, base) + "." + variant + ".proto"
}

var (
	testProtoRe = regexp.MustCompile(`(internal/testprotos/.*[.]proto)`)
	goPackageRe = regexp.MustCompile(`option go_package = "([^"]+)";`)
	extRe       = regexp.MustCompile(`_ext = ([0-9]+);`)
)

func generateOpaqueDotProto(repoRoot, tmpDir, relPath string) {
	// relPath is e.g. internal/testprotos/annotation/annotation.proto
	ignored := func(p string) bool {
		return strings.HasPrefix(p, "internal/testprotos/irregular")
	}
	inFn := filepath.Join(repoRoot, relPath)

	// create .hybrid.proto variant
	hybridFn := variantFn(relPath, "hybrid")
	outFn := filepath.Join(tmpDir, hybridFn)
	check(gsed(outFn, inFn, []func(line string) string{
		func(line string) string {
			if strings.HasPrefix(line, "package ") {
				return strings.ReplaceAll(line, "package ", "package hybrid.")
			}
			return line
		},
		func(line string) string {
			if testProtoPath := testProtoRe.FindString(line); testProtoPath != "" && !ignored(testProtoPath) {
				hybridFn := variantFn(testProtoPath, "hybrid")
				return strings.ReplaceAll(line, testProtoPath, hybridFn)
			}
			return line
		},
		func(line string) string {
			if matches := goPackageRe.FindStringSubmatch(line); matches != nil {
				goPkg := matches[1]
				hybridGoPkg := strings.TrimSuffix(goPkg, "/") + "/" + filepath.Base(goPkg) + "_hybrid"
				goPackage := `option go_package = "` + hybridGoPkg + `";` + "\n"
				if strings.HasPrefix(relPath, "internal/testprotos/test3/") {
					// The test3 testproto must remain on syntax = "proto3";
					// and therefore cannot use the editions-only api_level.
					return goPackage
				}
				return goPackage +
					`import "google/protobuf/go_features.proto";` + "\n" +
					`option features.(pb.go).api_level = API_HYBRID;`
			}
			return line
		},
	}...))

	// create .opaque.proto variant
	opaqueFn := variantFn(relPath, "opaque")
	outFn = filepath.Join(tmpDir, opaqueFn)
	check(gsed(outFn, inFn, []func(line string) string{
		func(line string) string {
			if strings.HasPrefix(line, "package ") {
				return strings.ReplaceAll(line, "package ", "package opaque.")
			}
			return line
		},
		func(line string) string {
			if testProtoPath := testProtoRe.FindString(line); testProtoPath != "" && !ignored(testProtoPath) {
				hybridFn := variantFn(testProtoPath, "opaque")
				return strings.ReplaceAll(line, testProtoPath, hybridFn)
			}
			return line
		},
		func(line string) string {
			if matches := goPackageRe.FindStringSubmatch(line); matches != nil {
				goPkg := matches[1]
				opaqueGoPkg := strings.TrimSuffix(goPkg, "/") + "/" + filepath.Base(goPkg) + "_opaque"
				goPackage := `option go_package = "` + opaqueGoPkg + `";` + "\n"
				if strings.HasPrefix(relPath, "internal/testprotos/test3/") {
					// The test3 testproto must remain on syntax = "proto3";
					// and therefore cannot use the editions-only api_level.
					return goPackage
				}
				return goPackage +
					`import "google/protobuf/go_features.proto";` + "\n" +
					`option features.(pb.go).api_level = API_OPAQUE;`
			}
			return line
		},
		func(line string) string {
			return strings.ReplaceAll(line, `full_name: ".goproto`, `full_name: ".opaque.goproto`)
		},
		func(line string) string {
			return strings.ReplaceAll(line, `type: ".goproto`, `type: ".opaque.goproto`)
		},
		func(line string) string {
			if matches := extRe.FindStringSubmatch(line); matches != nil {
				trimmed := strings.TrimSuffix(matches[0], ";")
				return strings.ReplaceAll(line, trimmed, trimmed+"0")
			}
			return line
		},
	}...))
}

func generateOpaqueTestprotos() {
	tmpDir, err := os.MkdirTemp(repoRoot, "tmp")
	check(err)
	defer os.RemoveAll(tmpDir)

	// Generate variants using the Hybrid and Opaque API for all local proto
	// files (except version-locked files).
	dirs := []struct {
		path     string
		pkgPaths map[string]string // mapping of .proto path to Go package path
		annotate map[string]bool   // .proto files to annotate
		exclude  map[string]bool   // .proto files to exclude from generation
	}{
		{path: "internal/testprotos/required"},
		{path: "internal/testprotos/testeditions"},
		{
			path: "internal/testprotos/test3",
			exclude: map[string]bool{
				"internal/testprotos/test3/test_extension.proto": true,
			},
		},
		{path: "internal/testprotos/enums"},
		{path: "internal/testprotos/textpbeditions"},
		{path: "internal/testprotos/messageset"},
		{
			path: "internal/testprotos/lazy",
			exclude: map[string]bool{
				"internal/testprotos/lazy/lazy_extension_normalized_wire_test.proto": true,
				"internal/testprotos/lazy/lazy_normalized_wire_test.proto":           true,
				"internal/testprotos/lazy/lazy_extension_test.proto":                 true,
			},
		},
	}
	excludeRx := regexp.MustCompile(`legacy/.*/`)
	for _, d := range dirs {
		srcDir := filepath.Join(repoRoot, filepath.FromSlash(d.path))
		filepath.Walk(srcDir, func(srcPath string, _ os.FileInfo, _ error) error {
			if !strings.HasSuffix(srcPath, ".proto") || excludeRx.MatchString(srcPath) {
				return nil
			}
			if strings.HasSuffix(srcPath, ".opaque.proto") || strings.HasSuffix(srcPath, ".hybrid.proto") {
				return nil
			}
			relPath, err := filepath.Rel(repoRoot, srcPath)
			check(err)

			if d.exclude[filepath.ToSlash(relPath)] {
				return nil
			}

			generateOpaqueDotProto(repoRoot, tmpDir, relPath)
			return nil
		})
	}

	syncOutput(repoRoot, tmpDir)
}

func generateEditionsDefaults() {
	dest := filepath.Join(repoRoot, "internal", "editiondefaults", "editions_defaults.binpb")
	srcDescriptorProto := filepath.Join(protoRoot, "src", "google", "protobuf", "descriptor.proto")
	srcGoFeatures := filepath.Join(repoRoot, "src", "google", "protobuf", "go_features.proto")
	// The enum in Go string formats to "EDITION_${EDITION}" but protoc expects
	// the flag in the form "${EDITION}". To work around this, we trim the
	// "EDITION_" prefix.
	minEdition := strings.TrimPrefix(fmt.Sprint(editionssupport.Minimum), "EDITION_")
	maxEdition := strings.TrimPrefix(fmt.Sprint(editionssupport.MaximumKnown), "EDITION_")
	cmd := exec.Command(
		"protoc",
		"--edition_defaults_out", dest,
		"--edition_defaults_minimum", minEdition,
		"--edition_defaults_maximum", maxEdition,
		"-I"+filepath.Join(protoRoot, "src"), "-I"+filepath.Join(repoRoot, "src"),
		srcDescriptorProto, srcGoFeatures,
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("executing: %v\n%s\n", strings.Join(cmd.Args, " "), out)
	}
	check(err)
}

func generateLocalProtos() {
	tmpDir, err := os.MkdirTemp(repoRoot, "tmp")
	check(err)
	defer os.RemoveAll(tmpDir)

	// Generate all local proto files (except version-locked files).
	dirs := []struct {
		path     string
		pkgPaths map[string]string // mapping of .proto path to Go package path
		annotate map[string]bool   // .proto files to annotate
		exclude  map[string]bool   // .proto files to exclude from generation
	}{{
		path: "cmd/protoc-gen-go/testdata",
		pkgPaths: map[string]string{
			"cmd/protoc-gen-go/testdata/nopackage/nopackage.proto": "google.golang.org/protobuf/cmd/protoc-gen-go/testdata/nopackage",
		},
		annotate: map[string]bool{"cmd/protoc-gen-go/testdata/annotations/annotations.proto": true},
	}, {
		path:    "internal/testprotos",
		exclude: map[string]bool{"internal/testprotos/irregular/irregular.proto": true},
	}, {
		path: "src/",
	}}
	excludeRx := regexp.MustCompile(`legacy/.*/`)
	for _, d := range dirs {
		subDirs := map[string]bool{}

		srcDir := filepath.Join(repoRoot, filepath.FromSlash(d.path))
		filepath.Walk(srcDir, func(srcPath string, _ os.FileInfo, _ error) error {
			if !strings.HasSuffix(srcPath, ".proto") || excludeRx.MatchString(srcPath) {
				return nil
			}
			relPath, err := filepath.Rel(repoRoot, srcPath)
			check(err)

			srcRelPath, err := filepath.Rel(srcDir, srcPath)
			check(err)
			subDirs[filepath.Dir(srcRelPath)] = true

			if d.exclude[filepath.ToSlash(relPath)] {
				return nil
			}

			opts := "module=" + modulePath
			for protoPath, goPkgPath := range d.pkgPaths {
				opts += fmt.Sprintf(",M%v=%v", protoPath, goPkgPath)
			}
			if d.annotate[filepath.ToSlash(relPath)] {
				opts += ",annotate_code"
			}
			if strings.HasPrefix(relPath, "internal/testprotos/test3/") {
				variant := strings.TrimPrefix(relPath, "internal/testprotos/test3/")
				if idx := strings.IndexByte(variant, '/'); idx > -1 {
					variant = variant[:idx]
				}
				switch variant {
				case "test3_hybrid":
					opts += fmt.Sprintf(",apilevelM%v=%v", relPath, "API_HYBRID")
				case "test3_opaque":
					opts += fmt.Sprintf(",apilevelM%v=%v", relPath, "API_OPAQUE")
				}
			}
			protoc("-I"+filepath.Join(repoRoot, "src"), "-I"+filepath.Join(protoRoot, "src"), "-I"+repoRoot, "--go_out="+opts+":"+tmpDir, filepath.Join(repoRoot, relPath))
			return nil
		})

		// For directories in testdata, generate a test that links in all
		// generated packages to ensure that it builds and initializes properly.
		// This is done because "go build ./..." does not build sub-packages
		// under testdata.
		if filepath.Base(d.path) == "testdata" {
			var imports []string
			for sd := range subDirs {
				imports = append(imports, fmt.Sprintf("_ %q", path.Join(modulePath, d.path, filepath.ToSlash(sd))))
			}
			sort.Strings(imports)

			s := strings.Join(append(generatedPreamble, []string{
				"package main",
				"",
				"import (" + strings.Join(imports, "\n") + ")",
			}...), "\n")
			b, err := format.Source([]byte(s))
			check(err)
			check(os.WriteFile(filepath.Join(tmpDir, filepath.FromSlash(d.path+"/gen_test.go")), b, 0664))
		}
	}

	syncOutput(repoRoot, tmpDir)
}

func generateRemoteProtos() {
	tmpDir, err := os.MkdirTemp(repoRoot, "tmp")
	check(err)
	defer os.RemoveAll(tmpDir)

	// Generate all remote proto files.
	files := []struct{ prefix, path, goPkgPath string }{
		// Well-known protos.
		{"src", "google/protobuf/any.proto", ""},
		{"src", "google/protobuf/api.proto", ""},
		{"src", "google/protobuf/duration.proto", ""},
		{"src", "google/protobuf/empty.proto", ""},
		{"src", "google/protobuf/field_mask.proto", ""},
		{"src", "google/protobuf/source_context.proto", ""},
		{"src", "google/protobuf/struct.proto", ""},
		{"src", "google/protobuf/timestamp.proto", ""},
		{"src", "google/protobuf/type.proto", ""},
		{"src", "google/protobuf/wrappers.proto", ""},

		// Compiler protos.
		{"src", "google/protobuf/compiler/plugin.proto", ""},
		{"src", "google/protobuf/descriptor.proto", ""},

		// Conformance protos.
		{"", "conformance/conformance.proto", "google.golang.org/protobuf/internal/testprotos/conformance;conformance"},
		{"src", "google/protobuf/test_messages_proto2.proto", "google.golang.org/protobuf/internal/testprotos/conformance;conformance"},
		{"src", "google/protobuf/test_messages_proto3.proto", "google.golang.org/protobuf/internal/testprotos/conformance;conformance"},
		{"src", "editions/golden/test_messages_proto2_editions.proto", "google.golang.org/protobuf/internal/testprotos/conformance/editionsmigration;editions"},
		{"src", "editions/golden/test_messages_proto3_editions.proto", "google.golang.org/protobuf/internal/testprotos/conformance/editionsmigration;editions"},
		{"", "conformance/test_protos/test_messages_edition2023.proto", "google.golang.org/protobuf/internal/testprotos/conformance/editions;editions"},

		// Benchmark protos.
		// TODO: The protobuf repo no longer includes benchmarks.
		//       CL removing them says they are superceded by google/fleetbench:
		//         https://github.com/protocolbuffers/protobuf/commit/83c499de86224538e5d59adc3d0fa7fdb45b2c72
		//       But that project's proto benchmark files are very different:
		//         https://github.com/google/fleetbench/tree/main/fleetbench/proto
		//       For now, commenting these out until benchmarks in this repo can be
		//       reconciled with new fleetbench stuff.
		//{"benchmarks", "benchmarks.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks;benchmarks"},
		//{"benchmarks", "datasets/google_message1/proto2/benchmark_message1_proto2.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message1/proto2;proto2"},
		//{"benchmarks", "datasets/google_message1/proto3/benchmark_message1_proto3.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message1/proto3;proto3"},
		//{"benchmarks", "datasets/google_message2/benchmark_message2.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message2;google_message2"},
		//{"benchmarks", "datasets/google_message3/benchmark_message3.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message3;google_message3"},
		//{"benchmarks", "datasets/google_message3/benchmark_message3_1.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message3;google_message3"},
		//{"benchmarks", "datasets/google_message3/benchmark_message3_2.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message3;google_message3"},
		//{"benchmarks", "datasets/google_message3/benchmark_message3_3.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message3;google_message3"},
		//{"benchmarks", "datasets/google_message3/benchmark_message3_4.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message3;google_message3"},
		//{"benchmarks", "datasets/google_message3/benchmark_message3_5.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message3;google_message3"},
		//{"benchmarks", "datasets/google_message3/benchmark_message3_6.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message3;google_message3"},
		//{"benchmarks", "datasets/google_message3/benchmark_message3_7.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message3;google_message3"},
		//{"benchmarks", "datasets/google_message3/benchmark_message3_8.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message3;google_message3"},
		//{"benchmarks", "datasets/google_message4/benchmark_message4.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message4;google_message4"},
		//{"benchmarks", "datasets/google_message4/benchmark_message4_1.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message4;google_message4"},
		//{"benchmarks", "datasets/google_message4/benchmark_message4_2.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message4;google_message4"},
		//{"benchmarks", "datasets/google_message4/benchmark_message4_3.proto", "google.golang.org/protobuf/internal/testprotos/benchmarks/datasets/google_message4;google_message4"},
	}

	opts := "module=" + modulePath
	for _, file := range files {
		if file.goPkgPath != "" {
			opts += fmt.Sprintf(",M%v=%v", file.path, file.goPkgPath)
		}
	}
	for _, f := range files {
		protoc("-I"+protoRoot, "-I"+filepath.Join(protoRoot, f.prefix), "--go_out="+opts+":"+tmpDir, f.path)
	}

	syncOutput(repoRoot, tmpDir)
}

func protoc(args ...string) {
	// TODO: Remove --experimental_allow_proto3_optional flag.
	cmd := exec.Command(
		"protoc",
		"--plugin=protoc-gen-go="+os.Args[0],
		"--experimental_allow_proto3_optional")
	if slices.ContainsFunc(args, func(s string) bool {
		return strings.Contains(s, "cmd/protoc-gen-go/testdata/") ||
			strings.Contains(s, "internal/testprotos/")
	}) {
		// Our testprotos use edition features of upcoming editions that protoc
		// has not yet declared support for:
		cmd.Args = append(cmd.Args, "--experimental_editions")
	}
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = append(os.Environ(), "RUN_AS_PROTOC_PLUGIN=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("executing: %v\n%s\n", strings.Join(cmd.Args, " "), out)
	}
	check(err)
}

// generateIdentifiers generates an internal package for descriptor.proto
// and well-known types.
func generateIdentifiers(gen *protogen.Plugin, file *protogen.File) {
	if file.Desc.Package() != "google.protobuf" && file.Desc.Package() != "pb" {
		return
	}

	importPath := modulePath + "/internal/genid"
	base := strings.TrimSuffix(path.Base(file.Desc.Path()), ".proto")
	g := gen.NewGeneratedFile(importPath+"/"+base+"_gen.go", protogen.GoImportPath(importPath))
	for _, s := range generatedPreamble {
		g.P(s)
	}
	g.P("package ", path.Base(importPath))
	g.P()

	g.P("const ", file.GoDescriptorIdent.GoName, " = ", strconv.Quote(file.Desc.Path()))
	g.P()

	var processEnums func([]*protogen.Enum)
	var processMessages func([]*protogen.Message)
	var processExtensions func([]*protogen.Extension)
	const protoreflectPackage = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")
	processEnums = func(enums []*protogen.Enum) {
		for _, enum := range enums {
			g.P("// Full and short names for ", enum.Desc.FullName(), ".")
			g.P("const (")
			g.P(enum.GoIdent.GoName, "_enum_fullname = ", strconv.Quote(string(enum.Desc.FullName())))
			g.P(enum.GoIdent.GoName, "_enum_name = ", strconv.Quote(string(enum.Desc.Name())))
			g.P(")")
			g.P()

			g.P("// Enum values for ", enum.Desc.FullName(), ".")
			g.P("const (")
			for _, v := range enum.Values {
				g.P(v.GoIdent.GoName, "_enum_value = ", v.Desc.Number())
			}
			g.P(")")
			g.P()
		}
	}
	processMessages = func(messages []*protogen.Message) {
		for _, message := range messages {
			g.P("// Names for ", message.Desc.FullName(), ".")
			g.P("const (")
			g.P(message.GoIdent.GoName, "_message_name ", protoreflectPackage.Ident("Name"), " = ", strconv.Quote(string(message.Desc.Name())))
			g.P(message.GoIdent.GoName, "_message_fullname ", protoreflectPackage.Ident("FullName"), " = ", strconv.Quote(string(message.Desc.FullName())))
			g.P(")")
			g.P()

			if len(message.Fields) > 0 {
				g.P("// Field names for ", message.Desc.FullName(), ".")
				g.P("const (")
				for _, field := range message.Fields {
					g.P(message.GoIdent.GoName, "_", field.GoName, "_field_name ", protoreflectPackage.Ident("Name"), " = ", strconv.Quote(string(field.Desc.Name())))
				}
				g.P()
				for _, field := range message.Fields {
					g.P(message.GoIdent.GoName, "_", field.GoName, "_field_fullname ", protoreflectPackage.Ident("FullName"), " = ", strconv.Quote(string(field.Desc.FullName())))
				}
				g.P(")")
				g.P()

				g.P("// Field numbers for ", message.Desc.FullName(), ".")
				g.P("const (")
				for _, field := range message.Fields {
					g.P(message.GoIdent.GoName, "_", field.GoName, "_field_number ", protoreflectPackage.Ident("FieldNumber"), " = ", field.Desc.Number())
				}
				g.P(")")
				g.P()
			}

			if len(message.Oneofs) > 0 {
				g.P("// Oneof names for ", message.Desc.FullName(), ".")
				g.P("const (")
				for _, oneof := range message.Oneofs {
					g.P(message.GoIdent.GoName, "_", oneof.GoName, "_oneof_name ", protoreflectPackage.Ident("Name"), " = ", strconv.Quote(string(oneof.Desc.Name())))
				}
				g.P()
				for _, oneof := range message.Oneofs {
					g.P(message.GoIdent.GoName, "_", oneof.GoName, "_oneof_fullname ", protoreflectPackage.Ident("FullName"), " = ", strconv.Quote(string(oneof.Desc.FullName())))
				}
				g.P(")")
				g.P()
			}

			processEnums(message.Enums)
			processMessages(message.Messages)
			processExtensions(message.Extensions)
		}
	}
	processExtensions = func(extensions []*protogen.Extension) {
		if len(extensions) == 0 {
			return
		}

		g.P("// Extension numbers")
		g.P("const (")
		for _, ext := range extensions {
			g.P(ext.Extendee.GoIdent.GoName, "_", ext.GoName, "_ext_number ", protoreflectPackage.Ident("FieldNumber"), " = ", ext.Desc.Number())
		}
		g.P(")")
	}
	processEnums(file.Enums)
	processMessages(file.Messages)
	processExtensions(file.Extensions)
}

// generateSourceContextStringer generates the implementation for the
// protoreflect.SourcePath.String method by using information present
// in the descriptor.proto.
func generateSourceContextStringer(gen *protogen.Plugin, file *protogen.File) {
	if file.Desc.Path() != "google/protobuf/descriptor.proto" {
		return
	}

	importPath := modulePath + "/reflect/protoreflect"
	g := gen.NewGeneratedFile(importPath+"/source_gen.go", protogen.GoImportPath(importPath))
	for _, s := range generatedPreamble {
		g.P(s)
	}
	g.P("package ", path.Base(importPath))
	g.P()

	var messages []*protogen.Message
	for _, message := range file.Messages {
		if message.Desc.Name() == "FileDescriptorProto" {
			messages = append(messages, message)
		}
	}
	seen := make(map[*protogen.Message]bool)

	for len(messages) > 0 {
		m := messages[0]
		messages = messages[1:]
		if seen[m] {
			continue
		}
		seen[m] = true

		g.P("func (p *SourcePath) append", m.GoIdent.GoName, "(b []byte) []byte {")
		g.P("if len(*p) == 0 { return b }")
		g.P("switch (*p)[0] {")
		for _, f := range m.Fields {
			g.P("case ", f.Desc.Number(), ":")
			var cardinality string
			switch {
			case f.Desc.IsMap():
				panic("maps are not supported")
			case f.Desc.IsList():
				cardinality = "Repeated"
			default:
				cardinality = "Singular"
			}
			nextAppender := "nil"
			if f.Message != nil {
				nextAppender = "(*SourcePath).append" + f.Message.GoIdent.GoName
				messages = append(messages, f.Message)
			}
			g.P("b = p.append", cardinality, "Field(b, ", strconv.Quote(string(f.Desc.Name())), ", ", nextAppender, ")")
		}
		g.P("}")
		g.P("return b")
		g.P("}")
		g.P()
	}
}

func syncOutput(dstDir, srcDir string) {
	filepath.Walk(srcDir, func(srcPath string, _ os.FileInfo, _ error) error {
		if !strings.HasSuffix(srcPath, ".go") &&
			!strings.HasSuffix(srcPath, ".meta") &&
			!strings.HasSuffix(srcPath, ".proto") {
			return nil
		}
		relPath, err := filepath.Rel(srcDir, srcPath)
		check(err)
		dstPath := filepath.Join(dstDir, relPath)

		if run {
			if copyFile(dstPath, srcPath) {
				fmt.Println("#", relPath)
			}
		} else {
			cmd := exec.Command("diff", dstPath, srcPath, "-N", "-u")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
		return nil
	})
}

func copyFile(dstPath, srcPath string) (changed bool) {
	src, err := os.ReadFile(srcPath)
	check(err)
	check(os.MkdirAll(filepath.Dir(dstPath), 0775))
	dst, _ := os.ReadFile(dstPath)
	if bytes.Equal(src, dst) {
		return false
	}
	check(os.WriteFile(dstPath, src, 0664))
	return true
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
