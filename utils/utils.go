package utils

import (
	"context"
	"fmt"
	"golady-tracker/config"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

// return a client connection for github action
func GithubClient() *github.Client {
  config, err := config.LoadConfig("./config")
  if err != nil {
    fmt.Print(err)
  }
  ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{ AccessToken: config.GithubAouthToken },
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

  return client
}