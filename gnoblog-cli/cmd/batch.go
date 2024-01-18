package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gnolang/gno/tm2/pkg/errors"
	"github.com/peterbourgon/ff/v3/ffcli"
	"os"
	"path/filepath"
)

type batchCfg struct {
	base *baseCliCfg
	root string
}

func newBatchCommand(baseCfg *baseCliCfg) *ffcli.Command {
	cfg := &batchCfg{
		base: baseCfg,
	}
	fs := flag.NewFlagSet("batch", flag.ExitOnError)

	cfg.registerFlags(fs)

	return &ffcli.Command{
		Name:       "batch",
		ShortUsage: "batch [flags]",
		LongHelp:   "Batch post multiple blog posts found in root dir",
		FlagSet:    fs,
		Exec:       cfg.exec,
	}
}

func (b *batchCfg) registerFlags(fs *flag.FlagSet) {
	fs.StringVar(&b.root,
		"root",
		"./",
		"root path of folder containing post dirs",
	)
}

func (b *batchCfg) exec(_ context.Context, _ []string) error {
	if b.root == "" {
		return errors.New("batch root folder cannot be empty")
	}

	found, err := findFilePaths(b.root)
	if err != nil {
		return err
	}

	if len(found) == 0 {
		return errors.New("no post READMEs found")
	}

	return nil
}

// findFilePaths gathers the file paths for specific file types
func findFilePaths(startPath string) ([]string, error) {
	filePaths := make([]string, 0)

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing file: %w", err)
		}

		// Check if the file is a dir
		if info.IsDir() {
			return nil
		}

		if info.Name() == "README.md" {
			filePaths = append(filePaths, path)
		}
		return nil
	}

	// Walk the directory root recursively
	if walkErr := filepath.Walk(startPath, walkFn); walkErr != nil {
		return nil, fmt.Errorf("unable to walk directory, %w", walkErr)
	}

	return filePaths, nil
}
