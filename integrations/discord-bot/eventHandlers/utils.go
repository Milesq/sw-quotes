package eventHandlers

import "github.com/bwmarrin/discordgo"

func createResponder(s *discordgo.Session, channelID string) func(string) {
	return func(content string) {
		s.ChannelMessageSend(channelID, content)
	}
}
