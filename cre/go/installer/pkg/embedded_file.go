package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type embeddedFile struct {
	name    string
	content string
}

func (f *embeddedFile) write(tmpDir string) error {
	fullName := strings.ReplaceAll(filepath.Join(tmpDir, f.name), "/", string(os.PathSeparator))
	dir := filepath.Dir(fullName)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s\n%w", dir, err)
	}

	if err := os.WriteFile(fullName, []byte(f.content), os.ModePerm); err != nil {
		return fmt.Errorf("failed to write file %s\n%w", fullName, err)
	}

	return nil
}
