package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/milesq/sw-quotes/integrations/discord-bot/eventHandlers"
)

var TOKEN string

func init() {
	flag.StringVar(&TOKEN, "t", "", "Discord Bot Token")
	flag.Parse()
}

func main() {
	bot, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	bot.AddHandler(eventHandlers.MessageHandler)

	err = bot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Close()
}
