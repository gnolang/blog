package main

import (
	"context"
	"time"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

func initGithubClient() *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: opts.githubToken})
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)
	return client
}

func githubFetchIssues(client *github.Client, opts *github.IssueListByRepoOptions, owner string, repository string) ([]*github.Issue, error) {
	issues, _, err := client.Issues.ListByRepo(context.Background(), owner, repository, opts)
	if err != nil {
		return nil, err
	}

	return issues, nil
}

func githubFetchCommits(client *github.Client, opts *github.CommitsListOptions, owner string, repository string) ([]*github.RepositoryCommit, error) {
	commits, _, err := client.Repositories.ListCommits(context.Background(), owner, repository, opts)
	if err != nil {
		return nil, err
	}

	return commits, nil
}

func githubFetchPullRequests(client *github.Client, opts *github.PullRequestListOptions, owner string, repository string) ([]*github.PullRequest, error) {
	pullRequests, _, err := client.PullRequests.List(context.Background(), owner, repository, opts)
	if err != nil {
		return nil, err
	}

	return pullRequests, nil
}

func filterPullRequestByTime(pullRequests []*github.PullRequest, since time.Time) []*github.PullRequest {
	var filtered []*github.PullRequest
	for _, pullRequest := range pullRequests {
		if pullRequest.UpdatedAt.After(since) {
			filtered = append(filtered, pullRequest)
		}
	}
	return filtered
}

func filterIssueNotPR(issues []*github.Issue) []*github.Issue {
	var filtered []*github.Issue
	for _, issue := range issues {
		if issue.PullRequestLinks == nil {
			filtered = append(filtered, issue)
		}
	}
	return filtered
}
