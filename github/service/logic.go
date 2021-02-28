package service

// implt of our service interface

import (
  "log"
  "io/ioutil"
  "encoding/json"

  // internal pkg
  "golady-tracker/github/models"
  "golady-tracker/github/utils"
)

type githubService struct {
}

func NewGithubService() GithubService {
  return &githubService{}
}

// our github issues struct
var githubIssues models.GithubIssues

func (g githubService) GetAllIssues() (models.GithubIssues, error) {
  responseGithub := utils.RequestGithubOptBasicAuth("orgs/optyce/issues", true)
  responseBody, err := ioutil.ReadAll(responseGithub.Body)
  if err != nil {
    return nil, err
  }
  if err := json.Unmarshal(responseBody, &githubIssues); err != nil {
    return nil, err
  }
  return githubIssues, nil
}

func (g githubService) GetIssuesByProject(projectName string) (models.GithubIssues, error) {
  return nil, nil
}

func (g githubService) CreateIssueByProject(issue, projectName string) (string, error) {
  return "", nil
}
