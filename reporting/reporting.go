package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/go-github/v50/github"
	"net/http"
	"os"
	"time"

	"github.com/peterbourgon/ff/v3/ffcli"
)

func DefaultOpts() Opts {
	return Opts{
		changelog:              true,
		backlog:                true,
		curation:               true,
		tips:                   true,
		format:                 "markdown",
		since:                  "",
		twitterToken:           "",
		githubToken:            "",
		help:                   false,
		httpClient:             &http.Client{},
		twitterSearchTweetsUrl: "https://api.twitter.com/2/tweets/search/recent?query=%23gnotips&tweet.fields=created_at&max_results=100&user.fields=username&expansions=author_id",
		awesomeGnoRepoUrl:      "https://api.github.com/repos/gnolang/awesome-gno/issues",
		outputPath:             "./output/",
	}
}

var opts = DefaultOpts()

func main() {
	err := runMain(os.Args[1:])
	if err == flag.ErrHelp {
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		os.Exit(1)
	}
}

// TODO: Verify if we can use a template engine to format the output: Like prebuild a report template with a function to format the data
func runMain(args []string) error {
	var root *ffcli.Command
	{
		globalFlags := flag.NewFlagSet("gno-reporting", flag.ExitOnError)
		globalFlags.BoolVar(&opts.changelog, "changelog", opts.changelog, "generate changelog")
		globalFlags.BoolVar(&opts.backlog, "backlog", opts.backlog, "generate backlog")
		globalFlags.BoolVar(&opts.curation, "curation", opts.curation, "generate curation")
		globalFlags.BoolVar(&opts.tips, "tips", opts.tips, "generate tips")
		globalFlags.StringVar(&opts.since, "since", opts.since, "since date RFC 3339 (ex: 2003-01-19T00:00:00.000Z)")
		globalFlags.StringVar(&opts.twitterToken, "twitter-token", opts.twitterToken, "twitter token")
		globalFlags.StringVar(&opts.githubToken, "github-token", opts.githubToken, "github token")
		globalFlags.BoolVar(&opts.help, "help", false, "show help")
		globalFlags.StringVar(&opts.format, "format", opts.format, "output format")
		globalFlags.StringVar(&opts.outputPath, "output-path", opts.outputPath, "output directory path")
		root = &ffcli.Command{
			ShortUsage: "reporting [flags]",
			FlagSet:    globalFlags,
			Exec: func(ctx context.Context, args []string) error {
				var err error
				since := time.Time{}
				if opts.help {
					return flag.ErrHelp
				}
				if opts.twitterToken == "" && opts.tips {
					return fmt.Errorf("twitter token is required to fetch tips")
				}
				if opts.githubToken == "" && (opts.curation || opts.backlog || opts.changelog) {
					return fmt.Errorf("github token is required to fetch curation, backlog or changelog")
				}
				if opts.since != "" {
					since, err = time.Parse("2006-01-02T15:04:05.000Z", opts.since)
					if err != nil {
						return err
					}
					if err != nil {
						return fmt.Errorf("invalid from date")
					}
				}
				githubClient := initGithubClient()
				outputFile, err := createOutputFile()
				err = fetchChangelog(githubClient, since, outputFile)
				if err != nil {
					return err
				}
				err = fetchBacklog(githubClient, since, outputFile)
				if err != nil {
					return err
				}
				err = fetchCuration(githubClient, since, outputFile)
				if err != nil {
					return err
				}
				err = fetchTips(outputFile)
				if err != nil {
					return err
				}
				return nil
			},
		}
	}
	return root.ParseAndRun(context.Background(), args)
}

// TODO: Fetch changelog recent contributors, new PR merged, new issues closed ...
func fetchChangelog(client *github.Client, since time.Time, outputFile *os.File) error {
	if !opts.changelog {
		return nil
	}

	issues, err := githubFetchIssues(client, &github.IssueListByRepoOptions{State: "closed", Since: since}, "gnolang", "gno")
	if err != nil {
		return err
	}
	issuesFiltered := filterIssueNotPR(issues)
	pullRequests, err := githubFetchPullRequests(client, &github.PullRequestListOptions{State: "closed"}, "gnolang", "gno")
	if err != nil {
		return err
	}
	pullRequestsFiltered := filterPullRequestByTime(pullRequests, since)
	commits, err := githubFetchCommits(client, &github.CommitsListOptions{Since: since}, "gnolang", "gno")
	if err != nil {
		return err
	}
	err = writeChangelog(issuesFiltered, pullRequestsFiltered, commits, outputFile)
	if err != nil {
		return err
	}
	return nil
}

func fetchBacklog(client *github.Client, since time.Time, outputFile *os.File) error {
	if !opts.backlog {
		return nil
	}
	issues, err := githubFetchIssues(client, &github.IssueListByRepoOptions{State: "open", Since: since}, "gnolang", "gno")
	if err != nil {
		return err
	}
	issuesFiltered := filterIssueNotPR(issues)
	pullRequests, err := githubFetchPullRequests(client, &github.PullRequestListOptions{State: "open"}, "gnolang", "gno")
	if err != nil {
		return err
	}
	pullRequestsFiltered := filterPullRequestByTime(pullRequests, since)
	if err != nil {
		return err
	}
	err = writeBacklog(issuesFiltered, pullRequestsFiltered, outputFile)
	if err != nil {
		return err
	}
	return nil
}

func fetchCuration(client *github.Client, since time.Time, outputFile *os.File) error {
	if !opts.curation {
		return nil
	}
	issues, err := githubFetchIssues(client, &github.IssueListByRepoOptions{State: "all", Since: since}, "gnolang", "awesome-gno")
	if err != nil {
		return err
	}
	issuesFiltered := filterIssueNotPR(issues)
	pullRequests, err := githubFetchPullRequests(client, &github.PullRequestListOptions{State: "all"}, "gnolang", "awesome-gno")
	if err != nil {
		return err
	}
	pullRequestsFiltered := filterPullRequestByTime(pullRequests, since)
	commits, err := githubFetchCommits(client, &github.CommitsListOptions{Since: since}, "gnolang", "awesome-gno")
	if err != nil {
		return err
	}
	err = writeCuration(issuesFiltered, pullRequestsFiltered, commits, outputFile)
	if err != nil {
		return err
	}
	return nil
}

func fetchTips(outputFile *os.File) error {
	if !opts.tips {
		return nil
	}
	ret := twitterFetchTips()
	err := writeTips(ret, outputFile)
	if err != nil {
		return err
	}
	return nil
}

type Opts struct {
	changelog              bool
	backlog                bool
	curation               bool
	tips                   bool
	since                  string
	twitterToken           string
	githubToken            string
	format                 string
	help                   bool
	httpClient             *http.Client
	twitterSearchTweetsUrl string
	awesomeGnoRepoUrl      string
	outputPath             string
}
