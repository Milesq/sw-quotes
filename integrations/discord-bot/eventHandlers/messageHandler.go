package eventHandlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

const prefix = "quote!"

func MessageHandler(s *discordgo.Session, msg *discordgo.MessageCreate) {
	respond := func(content string) {
		s.ChannelMessageSend(msg.ChannelID, content)
	}

	if s.State.User.ID == msg.Author.ID {
		return
	}

	if !strings.HasPrefix(msg.Content, prefix) {
		return
	}

	result, err := resolveQuery(msg.Content)

	if err != nil {
		respond("error")
	}

	respond(result.Srt)
}
