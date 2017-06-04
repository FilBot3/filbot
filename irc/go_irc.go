package main

import (
	"github.com/go-chat-bot/bot/irc"
	_ "github.com/predatorian3/filbot/hello_world"
	"os"
	"strings"
)

func main() {
	bot_config := &irc.Config{
		Server:   os.Getenv("IRC_SERVER"),
		Channels: strings.Split(os.Getenv("IRC_CHANNELS"), ","),
		User:     os.Getenv("IRC_USER"),
		Nick:     os.Getenv("IRC_NICK"),
		Password: os.Getenv("IRC_PASSWORD"),
		UseTLS:   true,
		Debug:    os.Getenv("DEBUG") != "",
	}
	irc.Run(bot_config)
}
