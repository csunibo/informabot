package main

import (
	"log"
	"os"

	"github.com/csunibo/informabot/bot"
)

const tokenKey = "TOKEN"

func main() {
	token, found := os.LookupEnv(tokenKey)
	if !found {
		log.Fatal("token not found. please set the TOKEN environment variable")
	}

	bot.StartInformaBot(token, false)
}
