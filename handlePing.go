package main

import (
	"github.com/bwmarrin/discordgo"
)

func handlePing(session *discordgo.Session, message *discordgo.MessageCreate) {
	_, _ = session.ChannelMessageSend(message.ChannelID, "Pong!")
}
