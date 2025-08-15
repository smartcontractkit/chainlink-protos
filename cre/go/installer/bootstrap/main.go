package main

import (
	"bytes"
	_ "embed"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

//go:embed embedded.go.tmpl
var embeddedTemplate string

func main() {
	sanitizedFiles := mustCollectFiles("capabilities", nil)
	sanitizedFiles = mustCollectFiles("sdk", sanitizedFiles)
	sanitizedFiles = mustCollectFiles("tools", sanitizedFiles)
	sanitizedFiles = mustCollectFiles("values", sanitizedFiles)

	results := bytes.Buffer{}
	err := template.Must(template.New("embedded").Parse(embeddedTemplate)).Execute(&results, sanitizedFiles)
	if err != nil {
		panic(err)
	}

	output := results.Bytes()
	formatted, formatErr := format.Source(output)
	if formatErr != nil {
		output = formatted
	}

	outputFile := filepath.Join("installer", "pkg", "embedded_gen.go")
	if err = os.WriteFile(outputFile, formatted, os.ModePerm); err != nil {
		panic(err)
	}

	if formatErr != nil {
		panic(formatErr)
	}
}

func mustCollectFiles(root string, sanitizedFiles []SanitizedFile) []SanitizedFile {
	if err := filepath.Walk(filepath.Join("..", root), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if ".proto" != filepath.Ext(path) {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(filepath.Join("..", root), path)
		if err != nil {
			return err
		}

		sanitizedFiles = append(sanitizedFiles, SanitizedFile{
			FileName: root + "/" + rel,
			VarName:  sanitize(rel),
			Content:  escapeBackticks(string(content)),
		})

		return nil
	}); err != nil {
		panic(err)
	}

	return sanitizedFiles
}

func sanitize(path string) string {
	path = strings.TrimSuffix(path, filepath.Ext(path))

	var out []rune
	capNext := false
	for i, r := range path {
		if i == 0 {
			out = append(out, unicode.ToLower(r))
			continue
		}
		if r == '_' || r == '-' || r == '.' || r == '/' || r == '\\' {
			capNext = true
			continue
		}
		if capNext {
			out = append(out, unicode.ToUpper(r))
			capNext = false
		} else {
			out = append(out, r)
		}
	}
	return string(out) + "Embedded"
}

func escapeBackticks(s string) string {
	return strings.ReplaceAll(s, "`", "` + \"`\" + `")
}

type SanitizedFile struct {
	FileName string
	VarName  string
	Content  string
}
