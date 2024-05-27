package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// extractTitle extracts H1 (Title) from the body of the Post
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

// removeTitle removes the first occurrence of the Title from the body of the Post
func removeTitle(body, title string) string {
	t := "# " + title + "\n"
	return strings.Replace(body, t, "", 1)
}

// removeWhitespace loops over a slice of tags (strings) and truncates each tag
// by removing whitespace
// ie, []string{"example spaced tag", "correct_tag"} -> []string{"examplespacedtag", "correct_tag"}
func removeWhitespace(tags []string) []string {
	t := make([]string, len(tags))

	for i, tag := range tags {
		t[i] = strings.Replace(tag, " ", "", -1)
	}

	return t
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

// askForConfirmation asks the user if they're sure that they want to post to a specific chain
func askForConfirmation(chainID string, keyName string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Are you sure you want to post to chain with id `%s` with key `%s`? (y/n):", chainID, keyName)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "y" {
			return true
		} else if input == "n" {
			return false
		} else {
			fmt.Println("Please enter 'y' or 'n'.")
		}
	}
}
