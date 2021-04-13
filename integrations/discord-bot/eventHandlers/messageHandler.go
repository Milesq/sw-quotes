package eventHandlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/milesq/sw-quotes/integrations/discord-bot/messages"
)

const prefix = "quote!"

func MessageHandler(s *discordgo.Session, msg *discordgo.MessageCreate) {
	respond := createResponder(s, msg.ChannelID)

	if s.State.User.ID == msg.Author.ID {
		return
	}

	if !strings.HasPrefix(msg.Content, prefix) {
		return
	}

	if msg.Content == prefix+"help" {
		respond(messages.HELP_MSG)
		return
	}

	result, err := resolveQuery(msg.Content)

	if err != nil {
		respond("error")
	}

	respond(result.Srt)
}
