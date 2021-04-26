package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v35/github"
	"golady-tracker/utils"
)

var client = utils.GithubClient()

func ReturnIssuesFromRepoName(repoOwner, repoName string) ([]*github.Issue, error) {
	fmt.Println(repoOwner, repoName)
	issues, _, err := client.Issues.ListByRepo(context.Background(), repoOwner, repoName, nil)
	if err != nil {
		fmt.Print(err)
		return issues, err
	}
	return issues, nil
}

/*
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
*/