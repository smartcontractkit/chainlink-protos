package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/smartcontractkit/chainlink-protos/cre/go/installer/pkg"
)

func main() {
	gen := &pkg.ProtocGen{Plugins: []pkg.Plugin{pkg.GoPlugin}}
	mustGenAndMove(gen, "values/v1/values.proto", "pb")
	mustGenAndMove(gen, "sdk/v1beta/sdk.proto", "")
	mustGenAndMove(gen, "tools/generator/v1beta/cre_metadata.proto", "")
}

func mustGenAndMove(gen *pkg.ProtocGen, file, innerPkg string) {
	if err := gen.GenerateFile(file, "."); err != nil {
		panic(err)
	}

	oldFile := strings.Replace(file, ".proto", ".pb.go", 1)

	var newLocation string
	if innerPkg == "" {
		newLocation = filepath.Join(filepath.Dir(filepath.Dir(oldFile)), filepath.Base(oldFile))
	} else {
		newLocation = filepath.Join(filepath.Dir(filepath.Dir(oldFile)), innerPkg, filepath.Base(oldFile))
	}

	if err := os.MkdirAll(filepath.Dir(newLocation), 0o755); err != nil {
		panic(err)
	}

	if err := os.Rename(oldFile, newLocation); err != nil {
		panic(err)
	}

	if err := os.RemoveAll(filepath.Dir(oldFile)); err != nil {
		panic(err)
	}
}
