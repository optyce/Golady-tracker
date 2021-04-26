package discord

import (
	"fmt"
	"golady-tracker/github"

	"github.com/Lukaesebrot/dgc"
)


type ProjectCtx struct {
	repoOwner, repoName string
}

func initProjectCtx(ctx *dgc.Ctx) ProjectCtx {
	arguments := ctx.Arguments

	projectCtx := ProjectCtx{
		repoOwner : arguments.Get(0).Raw(),
		repoName : arguments.Get(1).Raw(),
	}

	return projectCtx
}

func CommandGetIssuesProjectName(ctx *dgc.Ctx) {
	projectCtx := initProjectCtx(ctx)

	issues, err := github.ReturnIssuesFromRepoName(projectCtx.repoOwner, projectCtx.repoName)
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range issues {
		var res string
		if value.Body == nil {
			res = *value.Title
		} else {
			res = *value.Title + ": " + *value.Body
		}
		if err := ctx.RespondText(res); err != nil {
			fmt.Println(err)
		}
	}
}

func CommandGetIDIssuesProjectName(ctx *dgc.Ctx) {
	projectCtx := initProjectCtx(ctx)

	issuesID, err := github.ReturnAllIssuesIDFromRepoName(projectCtx.repoOwner, projectCtx.repoName)
	if err != nil {
		fmt.Println(err)
	}

	var res string
	for _, value := range issuesID {
		res = string(*value.ID) + ": " + *value.Body
	}
	if err := ctx.RespondText(res); err != nil {
		fmt.Println(err)
	}
}