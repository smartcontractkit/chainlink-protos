package pkg

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var values = Packages{
	Go:    "github.com/smartcontractkit/chainlink-protos/cre/go/values/pb",
	Proto: "values/v1/values.proto",
}

var sdk = Packages{
	Go:    "github.com/smartcontractkit/chainlink-protos/cre/go/sdk",
	Proto: "sdk/v1beta/sdk.proto",
}

var creMetadata = Packages{
	Go:    "github.com/smartcontractkit/chainlink-protos/cre/go/tools/generator",
	Proto: "tools/generator/v1beta/cre_metadata.proto",
}

type ProtocGen struct {
	ProtocHelper
	packageNames map[string]string
	sources      []string
	init         bool
	Plugins      []Plugin
	dir          *protosDir
}

// LinkPackage directly links a package and does not require ProtocHelper to be set
func (p *ProtocGen) LinkPackage(pkgs Packages) {
	if p.packageNames == nil {
		p.packageNames = make(map[string]string)
	}
	p.packageNames[pkgs.Proto] = pkgs.Go
}

func (p *ProtocGen) LinkCapabilities(config *CapabilityConfig) {
	for _, file := range config.FullProtoFiles() {
		goPkg := p.FullGoPackageName(config)
		p.LinkPackage(Packages{Go: goPkg, Proto: file})
	}
}

func (p *ProtocGen) AddSourceDirectories(sources ...string) {
	p.sources = append(p.sources, sources...)
}

// GenerateFile generates a single file using protoc with the provided plugins and sources.
// Calling this method directly does not require ProtocHelper to be set.
func (p *ProtocGen) GenerateFile(file, from string) error {
	if err := p.doInit(); err != nil {
		return err
	}

	var args []string
	for _, pkg := range p.sources {
		args = append(args, "-I", pkg)
	}

	for _, plugin := range p.Plugins {
		prefix := fmt.Sprintf("%s_", plugin.Name)
		if plugin.Path != "" {
			sep := string(filepath.Separator)

			upDir := ""
			if from != "." {
				upLen := len(strings.Split(from, string([]byte{filepath.Separator})))
				upDir = strings.Repeat(".."+sep, upLen)
			}

			args = append(args, fmt.Sprintf("--plugin=protoc-gen-%s=%s%s%sprotoc-gen-%s", plugin.Name, upDir, plugin.Path, sep, plugin.Name))
		}

		args = append(args, fmt.Sprintf("--%sout=.", prefix))
		args = append(args, fmt.Sprintf("--%sopt=paths=source_relative", prefix))

		for proto, goPkg := range p.packageNames {
			args = append(args, fmt.Sprintf("--%sopt=M%s=%s", prefix, proto, goPkg))
		}
	}

	args = append(args, file)
	out, err := run("protoc", from, args...)
	if err != nil {
		return fmt.Errorf("failed to run protoc: %v\n%s", err, out)
	}
	if out == "" {
		out = "No output"
	}
	fmt.Printf("Generated file %q\nexecution output:\n%s\n", file, out)

	return nil
}

func (p *ProtocGen) Generate(config *CapabilityConfig) error {
	return p.GenerateMany(map[string]*CapabilityConfig{".": config})
}

func (p *ProtocGen) GenerateMany(dirToConfig map[string]*CapabilityConfig) error {
	for _, config := range dirToConfig {
		p.LinkCapabilities(config)
	}

	fmt.Println("Generating capabilities")
	errMap := map[string]error{}
	for from, config := range dirToConfig {
		for _, file := range config.FullProtoFiles() {
			if err := p.GenerateFile(file, from); err != nil {
				errMap[file] = err
			}
		}
	}

	if len(errMap) > 0 {
		var errStrings []string
		for file, err := range errMap {
			if err != nil {
				errStrings = append(errStrings, fmt.Sprintf("file %s\n%v\n", file, err))
			}
		}

		return errors.New(strings.Join(errStrings, ""))
	}

	err := p.moveGeneratedFiles(dirToConfig)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProtocGen) moveGeneratedFiles(dirToConfig map[string]*CapabilityConfig) error {
	fmt.Println("Moving generated files to correct locations")
	for from, config := range dirToConfig {
		for i, file := range config.FullProtoFiles() {
			file = strings.Replace(file, ".proto", ".pb.go", 1)
			to := strings.Replace(config.Files[i], ".proto", ".pb.go", 1)
			if err := os.Rename(path.Join(from, file), path.Join(from, to)); err != nil {
				return fmt.Errorf("failed to move generated file %s: %w", file, err)
			}
		}

		if err := os.RemoveAll(path.Join(from, "capabilities")); err != nil {
			return fmt.Errorf("failed to remove capabilities directory %w", err)
		}
	}
	return nil
}

func (p *ProtocGen) doInit() error {
	if p.init {
		return nil
	}

	var err error
	p.dir, err = installProtosToTmpDir()
	if err != nil {
		return err
	}
	p.LinkPackage(values)
	p.LinkPackage(sdk)
	p.LinkPackage(creMetadata)

	p.AddSourceDirectories(p.dir.Dir)
	p.init = true
	return nil
}

func run(command string, path string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = path
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed running command\n%s\nfrom path: %s\nerror:%v\nexecution output:\n%s", cmd.String(), path, err, "\t"+strings.ReplaceAll(string(outputBytes), "\n", "\n\t"))
	}
	return strings.TrimSpace(string(outputBytes)), nil
}
