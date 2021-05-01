package eventHandlers

import (
	"fmt"
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

	command := strings.TrimPrefix(msg.Content, prefix)
	command = strings.TrimSpace(command)

	switch command {
	case "help":
		respond(messages.HELP_MSG)
		return
	case "scenes":
		respond(predefinedSceneInfo)
		return
	case "movies":
		result := "Available movies:\n"
		for _, movie := range availableMovies {
			result += fmt.Sprintln("- ", movie[0], ": `", movie[1], "`")
		}
		respond(result)
		return
	}

	result, err := resolveQuery(command)

	if err != nil {
		respond(err.Error())
		return
	}

	respond(result.Srt)
}
