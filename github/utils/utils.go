package utils

import (
  "fmt"
  "net/http"

  "golady-tracker/config"
)

// path /orgs/optyce/issues

func RequestGithubOptBasicAuth(pathName string, basicAuth bool) *http.Response {
  client := &http.Client{}
  githubPath := "https://api.github.com/" + pathName
  request, err := http.NewRequest("GET", githubPath, nil)
  if err != nil {
    fmt.Print(err)
    return nil
  }
  config, _ := config.LoadConfig("../config")
  if basicAuth != false {
    request.SetBasicAuth(config.GithubUsername, config.GithubTokenPassword)
    return nil
  }
  response, err := client.Do(request)
  if err != nil {
    fmt.Print(err)
    return nil
  }
  return response
}
