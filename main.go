package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	// internal pkg
	"golady-tracker/config"
	"golady-tracker/github/service"

	"github.com/bwmarrin/discordgo"
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

	// Register the messageCreate func as a callback for MessageCreate events.
	discord.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}
