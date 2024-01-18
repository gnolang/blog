package main

import (
	"fmt"
	"github.com/adrg/frontmatter"
	"io"
	"strings"
)

func parsePost(reader io.Reader) (*post, error) {
	var p post
	rest, err := frontmatter.MustParse(reader, &p)
	if err != nil {
		return nil, fmt.Errorf("invalid post: %w", err)
	}
	p.Body = string(rest)
	p.Tags = removeWhitespace(p.Tags)

	return &p, nil
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
