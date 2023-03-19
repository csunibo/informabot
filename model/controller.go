package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/csunibo/informabot/commands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (data MessageData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, data.Text)
	SendHTML(bot, msg)

	return ""
}

func (data HelpData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO HelpData: notimplemented, Got: %s\n", message.Text))
	SendHTML(bot, msg)

	return ""
}

func (data UpdateData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO UpdateData: notimplemented, Got: %s\n", message.Text))
	SendHTML(bot, msg)

	return ""
}

func (data LookingForData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO LookingForData: notimplemented, Got: %s\n", message.Text))
	SendHTML(bot, msg)

	return ""
}

func (data NotLookingForData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO NotLookingForData: notimplemented, Got: %s\n", message.Text))
	SendHTML(bot, msg)

	return ""
}

func (data YearlyData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	var chatTitle string = strings.ToLower(message.Chat.Title)

	// check if string starts with "Yearly"
	if strings.Contains(chatTitle, "primo") {
		return data.Command + "1"
	} else if strings.Contains(chatTitle, "secondo") {
		return data.Command + "2"
	} else if strings.Contains(chatTitle, "terzo") {
		return data.Command + "3"
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, data.NoYear)
		SendHTML(bot, msg)
	}

	return ""
}

func (data TodayLecturesData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	var todayTime time.Time = time.Now()
	var todayString string = todayTime.Format("2006-01-02")
	url := data.Url + fmt.Sprintf("&start=%s&end=%s", todayString, todayString)
	// TODO: print this url if bot debug mode is active

	var response string = commands.GetTimeTable(url)

	var msg tgbotapi.MessageConfig
	if response != "" {
		msg = tgbotapi.NewMessage(message.Chat.ID, data.Title+response)
	} else {
		msg = tgbotapi.NewMessage(message.Chat.ID, data.FallbackText)
	}
	SendHTML(bot, msg)

	return ""
}

func (data TomorrowLecturesData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	var todayTime time.Time = time.Now()
	var tomorrowTime time.Time = todayTime.AddDate(0, 0, 1)
	var tomorrowString string = tomorrowTime.Format("2006-01-02")
	url := data.Url + fmt.Sprintf("&start=%s&end=%s", tomorrowString, tomorrowString)

	var response string = commands.GetTimeTable(url)

	var msg tgbotapi.MessageConfig
	if response != "" {
		msg = tgbotapi.NewMessage(message.Chat.ID, data.Title+response)
	} else {
		msg = tgbotapi.NewMessage(message.Chat.ID, data.FallbackText)
	}
	SendHTML(bot, msg)

	return ""
}

func (data ListData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO ListData: notimplemented, Got: %s\n", message.Text))
	SendHTML(bot, msg)

	return ""
}

func (data CourseData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	emails := strings.Join(data.Professors, "@unibo.it\n ") + "@unibo.it\n"
	ternary_assignment := func(condition bool, true_value string) string {
		if condition {
			return true_value
		} else {
			return ""
		}
	}

	msg := tgbotapi.NewMessage(message.Chat.ID,
		ternary_assignment(data.Name != "", fmt.Sprintf("<b>%s</b>\n", data.Name))+
			ternary_assignment(data.Virtuale != "", fmt.Sprintf("<a href='https://virtuale.unibo.it/course/view.php?id=%s'>Virtuale</a>", data.Virtuale))+"\n"+
			ternary_assignment(data.Teams != "", fmt.Sprintf("<a href='https://teams.microsoft.com/l/meetup-join/19%%3ameeting_%s", data.Teams))+"%40thread.v2/0?context=%7b%22Tid%22%3a%22e99647dc-1b08-454a-bf8c-699181b389ab%22%2c%22Oid%22%3a%22080683d2-51aa-4842-aa73-291a43203f71%22%7d'>Videolezione</a>\n"+
			ternary_assignment(data.Website != "", fmt.Sprintf("<a href='https://www.unibo.it/it/didattica/insegnamenti/insegnamento/%s'>Sito</a>\n<a href='https://www.unibo.it/it/didattica/insegnamenti/insegnamento/%s/orariolezioni'>Orario</a>", data.Website, data.Website))+"\n"+
			ternary_assignment(data.Professors != nil, fmt.Sprintf("Professori:\n %s", emails))+
			ternary_assignment(data.Name != "", fmt.Sprintf("<a href='https://csunibo.github.io/%s/'>📚 Risorse: materiali, libri, prove</a>\n", ToKebabCase(data.Name)))+
			ternary_assignment(data.Name != "", fmt.Sprintf("<a href='https://github.com/csunibo/%s/'>📂 Repository GitHub delle risorse</a>\n", ToKebabCase(data.Name)))+
			ternary_assignment(data.Telegram != "", fmt.Sprintf("<a href='t.me/$%s'>👥 Gruppo Studenti</a>\n", data.Telegram)))
	SendHTML(bot, msg)

	return ""
}

func (data LuckData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO LuckData: notimplemented, Got: %s\n", message.Text))
	SendHTML(bot, msg)

	return ""
}

func (data InvalidData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO InvalidData: notimplemented, Got: %s\n", message.Text))
	SendHTML(bot, msg)

	return ""
}
