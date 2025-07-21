package main

import "time"

type Post struct {
	Title           string     // parsed out from body
	Tags            []string   `yaml:"tags"`
	Authors         []string   `yaml:"authors"`
	PublicationDate *time.Time `yaml:"publication_date"`
	Slug            string     `yaml:"slug"`
	Body            string
}
