package main

import (
	"fmt"
	"io"
	"time"

	"github.com/adrg/frontmatter"
)

type post struct {
	Title           string    `yaml:"title"`
	Tags            []string  `yaml:"tags"`
	Authors         []string  `yaml:"authors"`
	PublicationDate time.Time `yaml:"publication_date"`
	Slug            string    `yaml:"slug"`
	Body            string
}

func parsePost(reader io.Reader) (*post, error) {
	var p post
	rest, err := frontmatter.MustParse(reader, &p)
	if err != nil {
		return nil, fmt.Errorf("invalid post: %w", err)
	}
	p.Body = string(rest)

	return &p, nil
}
