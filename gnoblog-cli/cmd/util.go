package main

import (
	"bufio"
	"fmt"
	"github.com/adrg/frontmatter"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func parsePost(reader io.Reader) (*post, error) {
	var p post
	rest, err := frontmatter.MustParse(reader, &p)
	if err != nil {
		return nil, fmt.Errorf("invalid post frontmatter: %w", err)
	}

	p.Title, err = extractTitle(p.Body)
	p.Body = removeTitle(p.Body, p.Title)
	p.Tags = removeWhitespace(p.Tags)

	if p.PublicationDate == nil {
		now := time.Now()
		p.PublicationDate = &now
	}

	if err != nil {
		return nil, err
	}

	return &p, nil
}

// extractTitle extracts H1 (Title) from the body of the post
func extractTitle(body string) (string, error) {
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

// removeTitle removes the first occurrence of a H1 (Title) from the body of the post
func removeTitle(body, title string) string {
	t := "# " + title + "\n"
	return strings.Replace(body, t, "", 1)
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
