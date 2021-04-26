package main

import (
	"fmt"
	"github.com/Lukaesebrot/dgc"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"

	// internal pkg
	"golady-tracker/config"
	cmd "golady-tracker/discord"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		fmt.Println(err)
	}

	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a dgc router
	// NOTE: The dgc.Create function makes sure all the maps get initialized
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
		Name: "obj",

		// We also want the command to get triggered with the 'object' alias
		Aliases: []string{
			"object",
		},

		// These fields get displayed in the default help messages
		Description: "Responds with the injected custom object",
		Usage:       "obj",
		Example:     "obj",

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
		Handler: cmd.CommandGetIssuesProjectName,
	})

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	router.Initialize(discord)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

