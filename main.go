package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	cmd "golady-tracker/discord"
	"os"
	"os/signal"
	"syscall"

	// internal pkg
	"golady-tracker/config"
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

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	//  call our command
	cmd.CreateCommand(discord)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
