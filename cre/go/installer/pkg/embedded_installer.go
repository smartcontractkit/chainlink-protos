package pkg

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type protosDir struct {
	Dir string
}

func (p *protosDir) Close() error {
	return os.RemoveAll(p.Dir)
}

func installProtosToTmpDir() (*protosDir, error) {
	dir, err := os.MkdirTemp("", "cre-protos-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary directory for protos: %w", err)
	}

	for _, file := range allFiles {
		if err := file.write(dir); err != nil {
			return nil, err
		}
	}

	setupSignalHandler(dir)
	return &protosDir{Dir: dir}, nil
}

func setupSignalHandler(dir string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.RemoveAll(dir)
		os.Exit(1)
	}()
}
