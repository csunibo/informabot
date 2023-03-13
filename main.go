package main

import (
	"log"
	"os"

	"github.com/csunibo/informabot/bot"
	"github.com/csunibo/informabot/parse"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go <telegram-bot-token>")
	}
	bot.StartInformaBot(os.Args[1], false)
	parse.ParseActions()

}
