package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {

	// No need for the bot to process its own commands really
	if message.Author.ID == BotID {
		return
	}

	// Grab Command
	contentSplit := strings.Split(message.Content, ",")
	command := contentSplit[0]

	// Check and Dispatch!
	switch command {
	case ".ping":
		handlePing(session, message)
	case ".meetup":
		handleMeetup(session, message)
	default:
		return
	}
}
