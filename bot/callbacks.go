package bot

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/musianisamuele/telegram-bot-api"

	"github.com/csunibo/informabot/commands"
	"github.com/csunibo/informabot/model"
)

// Handle the callback for the lectures command (`/lezioni`)
// Parse the `callback_text` to check which operation it must to do:
// - If the string ends with "_today" or "_tomorrow" it returns the timetable
// - If the string just contains a "_y_<number>" it asks for today or tomorrow
// - Otherwise prints the course year of what timetable the user wants to see
func lecturesCallback(bot *tgbotapi.BotAPI, update *tgbotapi.Update, callback_text string) {
	var chatId = int64(update.CallbackQuery.Message.Chat.ID)
	var messageId = update.CallbackQuery.Message.MessageID

	if strings.HasSuffix(callback_text, "_today") || strings.HasSuffix(callback_text, "_tomorrow") {
		timeForLectures := time.Now()

		if strings.HasSuffix(callback_text, "_tomorrow") {
			timeForLectures = timeForLectures.AddDate(0, 0, 1)
		}

		yearRegex, err := regexp.Compile(`_y_(\d)_`)
		if err != nil {
			log.Printf("Error [yearRegex]: %s\n", err)
			return
		}

		year, err := strconv.Atoi(yearRegex.FindString(callback_text)[3:4])
		if err != nil {
			log.Printf("Error [convert to integer the year regex]: %s\n", err)
			return
		}

		cdlKey := callback_text[len("lectures_"):strings.Index(callback_text, "_y_")]

		cdl := model.Cdls[cdlKey]
		response, err := commands.GetTimeTable(cdl.Type, cdl.Name, cdl.Curriculum, year, timeForLectures)
		if err != nil {
			log.Printf("Error [GetTimeTable]: %s\n", err)
		}

		editConfig := tgbotapi.NewEditMessageText(chatId, messageId, response)
		editConfig.ParseMode = tgbotapi.ModeHTML
		_, err = bot.Send(editConfig)
		if err != nil {
			log.Printf("Error [bot.Send() for the NewEditMessageText]: %s\n", err)
		}
	} else if strings.Contains(callback_text, "_y_") {
		rows := make([][]tgbotapi.InlineKeyboardButton, 2)
		rows[0] = tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Oggi", fmt.Sprintf("%s_today", callback_text)))
		rows[1] = tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Domani", fmt.Sprintf("%s_tomorrow", callback_text)))

		keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)
		editConfig := tgbotapi.NewEditMessageReplyMarkup(chatId, messageId, keyboard)
		_, err := bot.Send(editConfig)
		if err != nil {
			log.Printf("Error [bot.Send() for the NewEditMessageReplyMarkup]: %s\n", err)
		}
	} else {
		cdlName := strings.TrimPrefix(callback_text, "lectures_")
		yearsNro := 3
		// Master degrees has a duration of only 2 years
		if strings.HasPrefix(callback_text, "lectures_lm") {
			yearsNro = 2
		}
		rows := make([][]tgbotapi.InlineKeyboardButton, yearsNro)

		i := 1
		for i <= yearsNro {
			buttonText := fmt.Sprintf("%s: %d^ anno", model.Cdls[cdlName].Course, i)
			buttonCallback := fmt.Sprintf("%s_y_%d", callback_text, i)
			row := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(buttonText, buttonCallback))
			rows[i-1] = row

			i++
		}
		keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)
		editConfig := tgbotapi.NewEditMessageReplyMarkup(chatId, messageId, keyboard)
		_, err := bot.Send(editConfig)
		if err != nil {
			log.Printf("Error [bot.Send() for the NewEditMessageReplyMarkup]: %s\n", err)
		}
	}
}
