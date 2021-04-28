package discord

import (
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
	"time"
)

func CreateCommand(discord *discordgo.Session) {
	router := dgc.Create(&dgc.Router{
		// We will allow '!' and 'example!' as the bot prefixes
		Prefixes: []string{
			"!",
			"example!",
		},

		IgnorePrefixCase: true,

		BotsAllowed: false,

		Commands: []*dgc.Command{},

		Middlewares: []dgc.Middleware{},

		PingHandler: func(ctx *dgc.Ctx) {
			ctx.RespondText("Pong!")
		},
	})

	router.RegisterCmd(&dgc.Command{
		// We want to use 'obj' as the primary name of the command
		Name: "getIssues",

		// We also want the command to get triggered with the 'object' alias
		Aliases: []string{
			"getIssues",
		},

		// These fields get displayed in the default help messages
		Description: "Return all issues from project owner and repo name :)",
		Usage:       "!getIssues owner name",
		Example:     "!getIssues optyce influx",

		// You can assign custom flags to a command to use them in middlewares
		Flags: []string{},

		// We want to ignore the command case
		IgnoreCase: true,

		// You may define sub commands in here
		SubCommands: []*dgc.Command{},

		// We want the user to be able to execute this command once in five seconds and the cleanup interval shpuld be one second
		RateLimiter: dgc.NewRateLimiter(5*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
			ctx.RespondText("You are being rate limited!")
		}),
		// Now we want to define the command handler
		Handler: CommandGetIssuesProjectName,
	})

	router.RegisterCmd(&dgc.Command{
		// We want to use 'obj' as the primary name of the command
		Name: "getIssuesID",

		// We also want the command to get triggered with the 'object' alias
		Aliases: []string{
			"getIssuesID",
		},

		// These fields get displayed in the default help messages
		Description: "Return all issues from project owner and repo name :)",
		Usage:       "!getIssuesID owner name",
		Example:     "!getIssues optyce influx",

		// You can assign custom flags to a command to use them in middlewares
		Flags: []string{},

		// We want to ignore the command case
		IgnoreCase: true,

		// You may define sub commands in here
		SubCommands: []*dgc.Command{},

		// We want the user to be able to execute this command once in five seconds and the cleanup interval shpuld be one second
		RateLimiter: dgc.NewRateLimiter(5*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
			ctx.RespondText("You are being rate limited!")
		}),
		// Now we want to define the command handler
		Handler: CommandGetIDIssuesProjectName,
	})

	router.Initialize(discord)
}
