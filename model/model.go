// In this file we define all the structs used to parse JSON files into Go
// structs

package model

import (
	tgbotapi "github.com/musianisamuele/telegram-bot-api"
)

type DataInterface interface {
	HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) CommandResponse
	HandleBotCallback(bot *tgbotapi.BotAPI, update *tgbotapi.Update, callback_text string)
	GetDescription() string
}

// GetActionFromType returns an Action struct with the right DataInterface,
// inferred from the commandType string
func GetActionFromType(name string, commandType string) Action {
	var data DataInterface
	switch commandType {
	case "message":
		data = MessageData{}
	case "help":
		data = HelpData{}
	case "issue":
		data = IssueData{}
	case "lookingFor":
		data = LookingForData{}
	case "notLookingFor":
		data = NotLookingForData{}
	case "buttonsLecture":
		data = Lectures{}
	case "buttonsRepresentatives":
		data = RepresentativesData{}
	case "list":
		data = ListData{}
	case "luck":
		data = LuckData{}
	default:
		data = InvalidData{}
	}

	return Action{
		Name: name,
		Type: commandType,
		Data: data,
	}
}

// SECTION GLOBAL JSON STRUCTS
type GroupsStruct = map[int64][]int64

type AutoReply struct {
	Text  string `json:"text"`
	Reply string `json:"reply"`
}

type SettingsStruct struct {
	MainGroupsIdentifiers []string `json:"mainGroupsIdentifiers"`
}

type Meme struct {
	Name string
	Text string
}

type Action struct {
	Name string
	Type string        `json:"type"`
	Data DataInterface `json:"data"`
}

type Mantainer struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

// config/teachings.json

type Teaching struct {
	Name       string   `json:"name"`
	Url        string   `json:"url"`
	Chat       string   `json:"chat"`
	Website    string   `json:"website"`
	Professors []string `json:"professors"`
}

// config/degrees.json

type YearStudyDiagram struct {
	Mandatory []string `json:"mandatory"`
	Electives []string `json:"electives"`
}

type Year struct {
	Year      int64            `json:"year"`
	Chat      string           `json:"chat"`
	Teachings YearStudyDiagram `json:"teachings"`
}

type Degree struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Years []Year `json:"years"`
	Chat  string `json:"chat"`
}

// timetables.json

type Curriculum struct {
	Name     string `json:"name"`
	Callback string `json:"callback"`
}

// Recognized by a callback string
type Timetable struct {
	Course       string `json:"course"`    // Course title
	Name         string `json:"name"`      // Course name
	Type         string `json:"type"`      // Type (laurea|magistrale|2cycle)
	Curriculum   string `json:"curricula"` // Curriculum
	Title        string `json:"title"`
	FallbackText string `json:"fallbackText"`
}

// SECTION ACTION STRUCTS DATA
type MessageData struct {
	Text        string `json:"text"`
	Description string `json:"description"`
}

type HelpData struct {
	Description string `json:"description"`
	Slashes     bool   `json:"slashes"`
}

type IssueData struct {
	Description string `json:"description"`
	Response    string `json:"response"`
	Fallback    string `json:"fallback"`
}

type LookingForData struct {
	Description  string `json:"description"`
	SingularText string `json:"singularText"`
	PluralText   string `json:"pluralText"`
	ChatError    string `json:"chatError"`
}

type NotLookingForData struct {
	Description   string `json:"description"`
	Text          string `json:"text"`
	ChatError     string `json:"chatError"`
	NotFoundError string `json:"notFoundError"`
}

type Lectures struct {
	Description  string `json:"description"`
	Title        string `json:"title"`
	FallbackText string `json:"fallbackText"`
}

type ListData struct {
	Description string     `json:"description"`
	Header      string     `json:"header"`
	Template    string     `json:"template"`
	Items       [][]string `json:"items"`
}

type LuckData struct {
	Description     string `json:"description"`
	NoLuckGroupText string `json:"noLuckGroupText"`
}

type RepresentativesData struct {
	Description  string `json:"description"`
	Title        string `json:"title"`
	FallbackText string `json:"fallbackText"`
}

type Representative struct {
	Course          string   `json:"course"`
	Representatives []string `json:"representatives"`
}

type InvalidData struct{}
