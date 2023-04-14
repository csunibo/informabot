package model

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/csunibo/informabot/commands"
	"github.com/csunibo/informabot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/exp/slices"
)

func (data MessageData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) CommandResponse {
	return makeResponseWithText(data.Text)
}

func (data HelpData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) CommandResponse {
	answer := ""
	for _, action := range Actions {
		if description := action.Data.GetDescription(); description != "" {
			answer += "/" + action.Name + " - " + description + "\n"
		}
	}

	return makeResponseWithText(answer)
}

func (data LookingForData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) CommandResponse {
	if (message.Chat.Type != "group" && message.Chat.Type != "supergroup") || slices.Contains(Settings.LookingForBlackList, message.Chat.ID) {
		log.Print("Error [LookingForData]: not a group or blacklisted")
		return makeResponseWithText(data.ChatError)
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
		return makeResponseWithText("Errore nel caricamento dei membri del gruppo")
	}

	var resultMsg string
	// Careful: additional arguments must be passed in the right order!
	if len(chatMembers) == 1 {
		resultMsg = fmt.Sprintf(data.SingularText, message.Chat.Title)
	} else {
		resultMsg = fmt.Sprintf(data.PluralText, len(chatMembers), message.Chat.Title)
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

	return makeResponseWithText(resultMsg)
}

func (data NotLookingForData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) CommandResponse {
	if (message.Chat.Type != "group" && message.Chat.Type != "supergroup") || slices.Contains(Settings.LookingForBlackList, message.Chat.ID) {
		log.Print("Error [NotLookingForData]: not a group or blacklisted")
		return makeResponseWithText(data.ChatError)
	} else if _, ok := Groups[message.Chat.ID]; !ok {
		log.Print("Info [NotLookingForData]: group empty, user not found")
		return makeResponseWithText(fmt.Sprintf(data.NotFoundError, message.Chat.Title))
	}

	var chatId = message.Chat.ID
	var senderId = message.From.ID
	var chatTitle = message.Chat.Title

	var msg string
	if idx := slices.Index(Groups[chatId], senderId); idx == -1 {
		log.Print("Info [NotLookingForData]: user not found in group")
		msg = fmt.Sprintf(data.NotFoundError, chatTitle)
	} else {
		Groups[chatId] = append(Groups[chatId][:idx], Groups[chatId][idx+1:]...)
		SaveGroups()
		msg = fmt.Sprintf(data.Text, chatTitle)
	}

	return makeResponseWithText(msg)
}

func (data YearlyData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) CommandResponse {
	var chatTitle string = strings.ToLower(message.Chat.Title)

	// check if string starts with "Yearly"
	if strings.Contains(chatTitle, "primo") {
		return makeResponseWithNextCommand(data.Command + "1")
	} else if strings.Contains(chatTitle, "secondo") {
		return makeResponseWithNextCommand(data.Command + "2")
	} else if strings.Contains(chatTitle, "terzo") {
		return makeResponseWithNextCommand(data.Command + "3")
	} else {
		return makeResponseWithText(data.NoYear)
	}
}

func (data TodayLecturesData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) CommandResponse {
	var todayTime time.Time = time.Now()
	var todayString string = todayTime.Format("2006-01-02")
	url := data.Url + fmt.Sprintf("&start=%s&end=%s", todayString, todayString)
	log.Printf("URL: %s\n", url)

	var response string = commands.GetTimeTable(url)

	var msg string
	if response != "" {
		msg = data.Title + response
	} else {
		msg = data.FallbackText
	}

	return makeResponseWithText(msg)
}

func (data TomorrowLecturesData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) CommandResponse {
	var todayTime time.Time = time.Now()
	var tomorrowTime time.Time = todayTime.AddDate(0, 0, 1)
	var tomorrowString string = tomorrowTime.Format("2006-01-02")
	url := data.Url + fmt.Sprintf("&start=%s&end=%s", tomorrowString, tomorrowString)

	var response string = commands.GetTimeTable(url)

	var msg string
	if response != "" {
		msg = data.Title + response
	} else {
		msg = data.FallbackText
	}

	return makeResponseWithText(msg)
}

func (data ListData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) CommandResponse {
	resultText := data.Header

	for _, item := range data.Items {
		itemInterface := make([]interface{}, len(item))
		for i, v := range item {
			itemInterface[i] = v
		}
		resultText += fmt.Sprintf(data.Template, itemInterface...)
	}

	return makeResponseWithText(resultText)
}

const BEGINNING_MONTH = time.September

func getCurrentAcademicYear() int {

	now := time.Now()
	year := now.Year()
	if now.Month() >= BEGINNING_MONTH {
		return year
	} else {
		return year - 1
	}
}

func (data CourseData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) CommandResponse {
	emails := strings.Join(data.Professors, "@unibo.it\n ") + "@unibo.it\n"
	ternary_assignment := func(condition bool, true_value string) string {
		if condition {
			return true_value
		} else {
			return ""
		}
	}

	currentAcademicYear := fmt.Sprint(getCurrentAcademicYear())
	msg := ternary_assignment(data.Name != "", fmt.Sprintf("<b>%s</b>\n", data.Name)) +
		ternary_assignment(data.Website != "", fmt.Sprintf("<a href='https://www.unibo.it/it/didattica/insegnamenti/insegnamento/%s/%s'>Sito</a>\n<a href='https://www.unibo.it/it/didattica/insegnamenti/insegnamento/%s/%s/orariolezioni'>Orario</a>", currentAcademicYear, data.Website, currentAcademicYear, data.Website)+"\n") +
		ternary_assignment(data.Professors != nil, fmt.Sprintf("Professori:\n %s", emails)) +
		ternary_assignment(data.Name != "", fmt.Sprintf("<a href='https://csunibo.github.io/%s/'>📚 Risorse: materiali, libri, prove</a>\n", utils.ToKebabCase(data.Name))) +
		ternary_assignment(data.Name != "", fmt.Sprintf("<a href='https://github.com/csunibo/%s/'>📂 Repository GitHub delle risorse</a>\n", utils.ToKebabCase(data.Name))) +
		ternary_assignment(data.Telegram != "", fmt.Sprintf("<a href='https://t.me/%s'>👥 Gruppo Studenti</a>\n", data.Telegram))

	return makeResponseWithText(msg)
}

func (data LuckData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) CommandResponse {
	var emojis = []string{"🎲", "🎯", "🏀", "⚽", "🎳", "🎰"}
	var noLuckGroups = []int64{-1563447632} // NOTE: better way to handle this?

	var canLuckGroup = true

	if slices.Index(noLuckGroups, message.Chat.ID) != -1 {
		canLuckGroup = false
	}

	var msg string
	if canLuckGroup {
		rand.NewSource(time.Now().Unix())
		emoji := emojis[rand.Intn(len(emojis))]

		msg = emoji
	} else {
		msg = data.NoLuckGroupText
	}

	return makeResponseWithText(msg)
}

func (data InvalidData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) CommandResponse {
	log.Printf("Probably a bug in the JSON action dictionary, got invalid data in command")

	return makeResponseWithText("Bot internal Error, contact developers")
}
