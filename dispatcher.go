package main

import (
	"github.com/bwmarrin/discordgo"
)

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {

	// No need for the bot to process its own commands really
	if message.Author.ID == BotID {
		return
	}

	// Grab Command

	// If Valid - DISPATCH!
}
