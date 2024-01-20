package main

import (
	"bufio"
	"fmt"
	"github.com/adrg/frontmatter"
	"github.com/gnolang/gno/tm2/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// known chainIDs
const (
	dev             = "dev"
	test3           = "test3"
	staging         = "staging"
	tendermint_test = "tendermint_test"
)

var knownChainIDs = []string{
	dev, test3, staging, tendermint_test,
}

func validateChainID(id string) error {
	for _, chainID := range knownChainIDs {
		if id == chainID {
			return nil
		}
	}
	return errors.New("chain ID does not match any known chainIDs")
}

func parsePost(reader io.Reader) (*post, error) {
	var p post
	rest, err := frontmatter.MustParse(reader, &p)
	if err != nil {
		return nil, fmt.Errorf("invalid post frontmatter: %w", err)
	}

	if p.PublicationDate == nil {
		now := time.Now()
		p.PublicationDate = &now
	}

	p.Body = string(rest)
	p.Tags = removeWhitespace(p.Tags)

	p.Title, err = parseTitle(p.Body)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// parseTitle extracts the first H1 from the body of the post
func parseTitle(body string) (string, error) {
	scanner := bufio.NewScanner(strings.NewReader(body))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "# ") {
			return line[2:], nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("no H1 header found")
}

// removeWhitespace loops over a slice of tags (strings) and truncates each tag
// by removing whitespace
// ie, []string{"example spaced tag"} -> []string{"examplespacedtag"}
func removeWhitespace(tags []string) []string {
	for i, tag := range tags {
		tags[i] = strings.Replace(tag, " ", "", -1)
	}

	return tags
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
