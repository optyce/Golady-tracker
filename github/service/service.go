package service

import (
  "golady-tracker/github/models"
)

type GithubService interface {
  GetAllIssues() (models.GithubIssues, error)
  GetIssuesByProject(projectName string) (models.GithubIssues, error)
  CreateIssueByProject(issue, projectName string) (string, error)
}
