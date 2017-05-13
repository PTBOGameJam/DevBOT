package main

import (
	"github.com/bwmarrin/discordgo"
)

func handlePing(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")
}
