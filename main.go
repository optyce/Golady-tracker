package main 

import (
	"fmt"

	// internal pkg
	"bug-tracker/config"

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
	}
	discord.Close()
}