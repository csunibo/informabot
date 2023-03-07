package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go <telegram-bot-token>")
	}
	bot, err := tgbotapi.NewBotAPI(os.Args[1])
	if err != nil {
		log.Fatalf("Error creating new bot: %s", err)
	}

	// debug = true -> print full requests and responses
	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Error getting updates: %s", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// TODO: fetch commands and replies list
		if update.Message.Command() == "start" {
			log.Printf("@%s: \t%s -> %s", update.Message.From.UserName, update.Message.Text, "hello world")
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "hello world")
			bot.Send(msg)
		} else {
			log.Printf("@%s: \t%s -> COMMAND NOT AVAILABLE", update.Message.From.UserName, update.Message.Text)
		}
	}
}
