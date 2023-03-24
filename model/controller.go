package model

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/csunibo/informabot/commands"
	"github.com/csunibo/informabot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/exp/slices"
)

func (data MessageData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, data.Text)
	utils.SendHTML(bot, msg)

	return ""
}

func (data HelpData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO HelpData: notimplemented, Got: %s\n", message.Text))
	utils.SendHTML(bot, msg)

	return ""
}

func (data UpdateData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO UpdateData: notimplemented, Got: %s\n", message.Text))
	utils.SendHTML(bot, msg)

	return ""
}

func (data LookingForData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	if (message.Chat.Type != "group" && message.Chat.Type != "supergroup") || slices.Contains(Settings.LookingForBlackList, message.Chat.ID) {
		utils.SendHTML(bot, tgbotapi.NewMessage(message.Chat.ID, data.ChatError))
		log.Print("Error [LookingForData]: not a group or blacklisted")
		return ""
	}

	var chatId = message.Chat.ID
	var senderID = message.From.ID

	log.Printf("LookingForData: %d, %d", chatId, senderID)
	if chatArray, ok := Groups[chatId]; ok {
		if !slices.Contains(chatArray, senderID) {
			Groups[chatId] = append(chatArray, senderID)
		}
	} else {
		Groups[chatId] = []int{senderID}
	}
	SaveGroups()

	chatMembers, err := utils.GetChatMembers(bot, message.Chat.ID, Groups[chatId])
	if err != nil {
		log.Printf("Error [LookingForData]: %s", err)
		return ""
	}

	var resultMsg string
	// NOTA: c'è una dipendenza molto forte con il json del testo qui.
	if len(chatMembers) == 1 {
		resultMsg = fmt.Sprintf(data.SingularText, message.Chat.Title)
	} else {
		resultMsg = fmt.Sprintf(data.PluralText, message.Chat.Title, len(chatMembers))
	}

	for _, member := range chatMembers {
		userLastName := ""
		if member.User.LastName != "" {
			userLastName = " " + member.User.LastName
		}
		resultMsg += fmt.Sprintf("👤 <a href='tg://user?id=%d'>%s%s</a>\n",
			member.User.ID,
			member.User.FirstName,
			userLastName)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, resultMsg)
	utils.SendHTML(bot, msg)

	return ""
}

func (data NotLookingForData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	if (message.Chat.Type != "group" && message.Chat.Type != "supergroup") || slices.Contains(Settings.LookingForBlackList, message.Chat.ID) {
		utils.SendHTML(bot, tgbotapi.NewMessage(message.Chat.ID, data.ChatError))
		log.Print("Error [NotLookingForData]: not a group or blacklisted")
		return ""
	} else if _, ok := Groups[message.Chat.ID]; !ok {
		log.Print("Info [NotLookingForData]: group empty, user not found")
		utils.SendHTML(bot, tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(data.NotFoundError, chatTitle)))
		return ""
	}

	var chatId = message.Chat.ID
	var senderId = message.From.ID
	var chatTitle = message.Chat.Title

	var msg tgbotapi.MessageConfig

	if idx := slices.Index(Groups[chatId], senderId); idx == -1 {
		log.Print("Info [NotLookingForData]: user not found in group")
		msg = tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(data.NotFoundError, chatTitle))
	} else {
		Groups[chatId] = append(Groups[chatId][:idx], Groups[chatId][idx+1:]...)
		SaveGroups()
		msg = tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(data.Text, chatTitle))
	}

	utils.SendHTML(bot, msg)

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
		utils.SendHTML(bot, msg)
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
	utils.SendHTML(bot, msg)

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
	utils.SendHTML(bot, msg)

	return ""
}

func (data ListData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO ListData: notimplemented, Got: %s\n", message.Text))
	utils.SendHTML(bot, msg)

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
			ternary_assignment(data.Name != "", fmt.Sprintf("<a href='https://csunibo.github.io/%s/'>📚 Risorse: materiali, libri, prove</a>\n", utils.ToKebabCase(data.Name)))+
			ternary_assignment(data.Name != "", fmt.Sprintf("<a href='https://github.com/csunibo/%s/'>📂 Repository GitHub delle risorse</a>\n", utils.ToKebabCase(data.Name)))+
			ternary_assignment(data.Telegram != "", fmt.Sprintf("<a href='t.me/$%s'>👥 Gruppo Studenti</a>\n", data.Telegram)))
	utils.SendHTML(bot, msg)

	return ""
}

func (data LuckData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO LuckData: notimplemented, Got: %s\n", message.Text))
	utils.SendHTML(bot, msg)

	return ""
}

func (data InvalidData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO InvalidData: notimplemented, Got: %s\n", message.Text))
	utils.SendHTML(bot, msg)

	return ""
}
