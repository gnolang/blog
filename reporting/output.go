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

func writeChangelog(issues []*github.Issue, outputFile *os.File) error {
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
	result := fmt.Sprintf("# Changelog ⚙️\n\nThere is **%d new closed issues** in gnolang/gno since %s\n\n%s", len(issues), opts.since, markdownTable)

	_, err = outputFile.WriteString(result)
	if err != nil {
		return err
	}
	return nil
}

func writeBacklog(issues []*github.Issue, outputFile *os.File) error {
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
	result := fmt.Sprintf("# Backlog 💡\n\nThere is **%d new open issues** in gnolang/gno since %s\n\n%s", len(issues), opts.since, markdownTable)

	_, err = outputFile.WriteString(result)
	if err != nil {
		return err
	}
	return nil
}

func writeCuration(issues []*github.Issue, pullRequests []*github.PullRequest, commits []*github.RepositoryCommit, outputFile *os.File) error {
	// Format at Markdown format
	var issuesRows [][]string
	for _, issue := range issues {
		issuesRows = append(issuesRows, []string{issue.GetTitle(), issue.GetURL(), issue.GetUser().GetLogin()})
	}
	issuesTable, err := markdown.NewTableFormatterBuilder().
		Build("Title", "Link to issue", "Author").
		Format(issuesRows)
	if err != nil {
		return err
	}
	result := fmt.Sprintf("# Curation 📚\n\n## New issues\nThere is **%d new issues** in gnolang/awesome-gno since %s\n\n%s", len(issues), opts.since, issuesTable)

	var pullRequestRows [][]string
	for _, pr := range pullRequests {
		pullRequestRows = append(pullRequestRows, []string{pr.GetTitle(), pr.GetURL(), pr.GetUser().GetLogin()})
	}
	pullRequestsTable, err := markdown.NewTableFormatterBuilder().
		Build("Title", "Link to PR", "Author").
		Format(pullRequestRows)
	if err != nil {
		return err
	}
	result += fmt.Sprintf("\n\n## New PR\nThere is **%d new PR** in gnolang/awesome-gno since %s\n\n%s", len(pullRequests), opts.since, pullRequestsTable)

	result += fmt.Sprintf("\n\n## New commits\nThere is **%d new commits** in gnolang/awesome-gno since %s\n\n", len(commits), opts.since)

	_, err = outputFile.WriteString(result)
	if err != nil {
		return err
	}
	return nil
}

func writeTips(data string, outputFile *os.File) error {
	// Format at Markdown format
	var rows [][]string
	var tweets TweetSearch
	authors := make(map[string]string)

	_ = json.Unmarshal([]byte(data), &tweets)
	for _, user := range tweets.Includes.Users {
		authors[user.Id] = user.Username
	}
	for _, tweet := range tweets.Data {
		rows = append(rows, []string{authors[tweet.AuthorId], strings.Replace(tweet.Text, "\n", " ", -1), tweet.CreatedAt})
	}

	//Maybe build our own table formatter
	table, err := markdown.NewTableFormatterBuilder().
		Build("Author", "Text", "Created at").
		Format(rows)
	if err != nil {
		return err
	}
	result := fmt.Sprintf("# Tips 🐦\n\n## New Tweets with #gnotips\nThere is **%d new tweet** about gnotips since %s\n\n%s", tweets.Meta.ResultCount, opts.since, table)

	_, err = outputFile.WriteString(result)
	if err != nil {
		return err
	}
	return nil
}
