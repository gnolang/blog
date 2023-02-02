package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
	"github.com/google/go-github/v50/github"
	"os"
	"strings"
)

func jsonFormat(data string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(data), "", "\t")
	if err != nil {
		return data
	}
	return out.String()
}

func createOutputFile() (*os.File, error) {
	if _, err := os.Stat(opts.outputPath); os.IsNotExist(err) {
		err = os.MkdirAll(opts.outputPath, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	outputFile, err := os.Create(opts.outputPath + "report.md")
	if err != nil {
		return nil, err
	}
	return outputFile, nil
}

func writeChangelog(data string, outputFile *os.File) error {
	if opts.format == "json" {
		data = jsonFormat(data)
	}
	_, err := outputFile.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

func writeBacklog(data string, outputFile *os.File) error {
	if opts.format == "json" {
		data = jsonFormat(data)
	}
	_, err := outputFile.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

func writeCuration(issues []*github.Issue, outputFile *os.File) error {
	var issuesTable [][]string
	for _, issue := range issues {
		issuesTable = append(issuesTable, []string{issue.GetTitle(), issue.GetURL(), issue.GetUser().GetLogin()})
	}

	markdownTable, err := markdown.NewTableFormatterBuilder().
		Build("Title", "Link to Body", "Assignee").
		Format(issuesTable)
	if err != nil {
		return err
	}
	result := fmt.Sprintf("# Curation\n\nThere is **%d new issues** in gno/awesome-gno since %s\n\n%s", len(issues), opts.since, markdownTable)

	_, err = outputFile.WriteString(result)
	if err != nil {
		return err
	}
	return nil
}

func writeTips(data string, outputFile *os.File) error {
	// Format at Markdown format
	var table [][]string
	var tweets TweetSearch
	authors := make(map[string]string)

	_ = json.Unmarshal([]byte(data), &tweets)
	for _, user := range tweets.Includes.Users {
		authors[user.Id] = user.Username
	}
	for _, tweet := range tweets.Data {
		table = append(table, []string{authors[tweet.AuthorId], strings.Replace(tweet.Text, "\n", " ", -1), tweet.CreatedAt})
	}

	//Maybe build our own table formatter
	markdownTable, err := markdown.NewTableFormatterBuilder().
		Build("Author", "Text", "Created at").
		Format(table)
	if err != nil {
		return err
	}
	result := fmt.Sprintf("# Tips\n\nThere is **%d new tweet** about gnotips since %s\n\n%s", tweets.Meta.ResultCount, opts.since, markdownTable)

	_, err = outputFile.WriteString(result)
	if err != nil {
		return err
	}
	return nil
}
