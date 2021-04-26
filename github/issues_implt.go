package github

import (
	"context"
	"github.com/google/go-github/v35/github"
	"golady-tracker/utils"
)

var client = utils.GithubClient()

func ReturnIssuesFromRepoName(repoOwner, repoName string) (interface{}, error) {
	issues, _, err := client.Issues.ListByRepo(context.Background(), repoOwner, repoName, nil)
	if err != nil {
		return issues, err
	}
	return issues, nil
}

func CreateIssuesFromRepoName(repoOwner, repoName, title string) (string, error) {
	issuesRequest := github.IssueRequest{
		Title: title,
	}
	createIssue, _, err := client.Issues.Create(repoOwner, repoName, *issuesRequest)
	if err != nil {
		return "", nil
	}
	return "Issues has been created", nil
}