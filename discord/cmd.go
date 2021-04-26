package discord

import (
	"encoding/json"
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

	structToString, err := json.Marshal(issues)
	if err != nil {
		fmt.Print(err)
	}
	ctx.RespondText(string(structToString))
}