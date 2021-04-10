package eventHandlers

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.State.User.ID == m.Author.ID {
		return
	}

	s.ChannelMessageSend(m.ChannelID, "Hello "+m.Author.Username)

	time.Sleep(time.Second * 2)

	s.ChannelMessageSend(m.ChannelID, "I'm basic discord bot written in GoLang btw")
}
