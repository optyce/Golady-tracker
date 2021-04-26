package github

import (
	"context"
	"fmt"
	"golady-tracker/utils"
)

func ReturnRepositoriesFromUsername(username string) interface{} {
	client := utils.GithubClient()
	repos, _, err := client.Repositories.List(context.Background(), username, nil)
	if err != nil {
		fmt.Print(err)
	}
	return repos
}




