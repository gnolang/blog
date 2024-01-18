package main

import "time"

type post struct {
	Title           string    `yaml:"title"`
	Tags            []string  `yaml:"tags"`
	Authors         []string  `yaml:"authors"`
	PublicationDate time.Time `yaml:"publication_date"`
	Slug            string    `yaml:"slug"`
	Body            string
}
