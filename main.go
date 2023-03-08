package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go <telegram-bot-token>")
	}
	StartInformaBot(os.Args[1], false)

}
