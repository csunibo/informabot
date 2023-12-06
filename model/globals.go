// This file contains all the global variables of the bot that are initialized
// with the start of the bot.
//
// This file should be here because it had circular imports with the Model (bot
// imported model, which imported bot in order to access the global variables,
// especially for the settings)

package model

import (
	"log"
)

var (
	Autoreplies     []AutoReply
	Actions         []Action
	Degrees         map[string]Degree
	MemeList        []Meme
	Settings        SettingsStruct
	Teachings       map[string]Teaching
	Groups          GroupsStruct
	Timetables      map[string]Timetable
	Mantainers      []Mantainer
	Representatives map[string]Representative
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

	Teachings, err = ParseTeachings()
	if err != nil {
		log.Fatalf(err.Error())
	}

	Degrees, err = ParseDegrees()
	if err != nil {
		log.Fatalf(err.Error())
	}

	Settings, err = ParseSettings()
	if err != nil {
		log.Fatalf("Error reading settings.json file: %s", err.Error())
	}

	MemeList, err = ParseMemeList()
	if err != nil {
		log.Fatalf("Error reading memes.json file: %s", err.Error())
	}

	Groups, err = ParseOrCreateGroups()
	if err != nil {
		log.Fatalf("Error reading or creating groups.json file: %s", err.Error())
	}

	Timetables, err = ParseTimetables()
	if err != nil {
		log.Fatalf(err.Error())
	}

	Mantainers, err = ParseMantainers()
	if err != nil {
		log.Fatalf("Error parsing mantainers.json file: %s", err.Error())
	}

	Representatives, err = ParseRepresentatives()
	if err != nil {
		log.Fatalf("Error parsing representatives.json file: %s", err.Error())
	}
}
