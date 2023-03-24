package model

import (
	"log"
)

// This file contains all the global variables of the bot, that are initialized
// with the start of the bot, this file should be here because it had circular
// imports with the Model (bot imported model, which imported bot in order to access
// the global variables (expecially for the settings))

var (
	Autoreplies []AutoReply
	Actions     []Action
	MemeList    []Meme
	Settings    SettingsStruct
)

func InitGlobals() {
	var err error
	Autoreplies, err = ParseAutoReplies()
	if err != nil {
		log.Fatalf("Error reading autoreply.json file: %s", err.Error())
	}

	Actions, err = ParseActions()
	if err != nil {
		log.Fatalf("Error reading actions.json file: %s", err.Error())
	}

	Settings, err = ParseSettings()
	if err != nil {
		log.Fatalf("Error reading settings.json file: %s", err.Error())
	}

	MemeList, err = ParseMemeList()
	if err != nil {
		log.Fatalf("Error reading memes.json file: %s", err.Error())
	}

}
