package discord

import (
	"fmt"
	"golady-tracker/github"

	"github.com/Lukaesebrot/dgc"
)

func CommandGetIssuesProjectName(ctx *dgc.Ctx) {
	arguments := ctx.Arguments

	repoOwner := arguments.Get(0).Raw()
	repoName := arguments.Get(1).Raw()

	issues, err := github.ReturnIssuesFromRepoName(repoOwner, repoName)
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range issues {
		var res string
		if value.Body == nil {
			res = *value.Title
		} else {
			res = *value.Title + ": " + *value.Body }
		if err := ctx.RespondText(res); err != nil {
			fmt.Println(err)
		}
	}
}