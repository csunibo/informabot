package model

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type DataInterface interface {
	// Returns another command to be executed, or emtpy string if no command is to be executed
	// NOTE: we assume that everything returned by this function is a valid command
	HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string
}

func GetActionFromType(name string, commandType string) Action {
	var data DataInterface
	switch commandType {
	case "message":
		data = MessageData{}
	case "help":
		data = HelpData{}
	case "update":
		data = UpdateData{}
	case "lookingFor":
		data = LookingForData{}
	case "notLookingFor":
		data = NotLookingForData{}
	case "yearly":
		data = YearlyData{}
	case "todayLectures":
		data = TodayLecturesData{}
	case "tomorrowLectures":
		data = TomorrowLecturesData{}
	case "list":
		data = ListData{}
	case "course":
		data = CourseData{}
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

type AutoReply struct {
	Text  string `json:"text"`
	Reply string `json:"reply"`
}

type Action struct {
	Name string
	Type string        `json:"type"`
	Data DataInterface `json:"data"`
}

type MessageData struct {
	Text string `json:"text"`
}

type HelpData struct {
	Description string `json:"description"`
}

type UpdateData struct {
	Description string `json:"description"`
	NoYear      string `json:"noYear"`
	NoMod       string `json:"noMod"`
	Started     string `json:"started"`
	Ended       string `json:"ended"`
	Failed      string `json:"failed"`
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

type YearlyData struct {
	Description string `json:"description"`
	Command     string `json:"command"`
	NoYear      string `json:"noYear"`
}

type TodayLecturesData struct {
	Description  string `json:"description"`
	Url          string `json:"url"`
	Title        string `json:"title"`
	FallbackText string `json:"fallbackText"`
}

type TomorrowLecturesData TodayLecturesData

type ListData struct {
	Description string     `json:"description"`
	Header      string     `json:"header"`
	Template    string     `json:"template"`
	Items       [][]string `json:"items"`
}

type CourseData struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Virtuale    string   `json:"virtuale"`
	Teams       string   `json:"teams"`
	Website     string   `json:"website"`
	Professors  []string `json:"professors"`
	Telegram    string   `json:"telegram"`
}

type LuckData struct {
	Description string `json:"description"`
}

type InvalidData struct{}

func (data MessageData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, data.Text)
	bot.Send(msg)

	return ""
}

func (data HelpData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO HelpData: notimplemented, Got: %s\n", message.Text))
	bot.Send(msg)

	return ""
}

func (data UpdateData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO UpdateData: notimplemented, Got: %s\n", message.Text))
	bot.Send(msg)

	return ""
}

func (data LookingForData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO LookingForData: notimplemented, Got: %s\n", message.Text))
	bot.Send(msg)

	return ""
}

func (data NotLookingForData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO NotLookingForData: notimplemented, Got: %s\n", message.Text))
	bot.Send(msg)

	return ""
}

func (data YearlyData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO YearlyData: notimplemented, Got: %s\n", message.Text))
	bot.Send(msg)

	return ""
}

func (data TodayLecturesData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO TodayLecturesData: notimplemented, Got: %s\n", message.Text))
	bot.Send(msg)

	return ""
}

func (data TomorrowLecturesData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO TomorrowLecturesData: notimplemented, Got: %s\n", message.Text))
	bot.Send(msg)

	return ""
}

func (data ListData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO ListData: notimplemented, Got: %s\n", message.Text))
	bot.Send(msg)

	return ""
}

func (data CourseData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO CourseData: notimplemented, Got: %s\n", message.Text))
	bot.Send(msg)

	return ""
}

func (data LuckData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO LuckData: notimplemented, Got: %s\n", message.Text))
	bot.Send(msg)

	return ""
}

func (data InvalidData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO InvalidData: notimplemented, Got: %s\n", message.Text))
	bot.Send(msg)

	return ""
}
