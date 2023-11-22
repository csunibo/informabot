package bot

import (
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

		timetableKey := callback_text[len("lectures_"):strings.Index(callback_text, "_y_")]

		timetable := model.Timetables[timetableKey]
		response, err := commands.GetTimeTable(timetable.Type, timetable.Name, timetable.Curriculum, timetable.Url, year, timeForLectures)
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
		rows := model.ChooseTimetableDay(callback_text)
		keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)
		editConfig := tgbotapi.NewEditMessageReplyMarkup(chatId, messageId, keyboard)
		_, err := bot.Send(editConfig)
		if err != nil {
			log.Printf("Error [bot.Send() for the NewEditMessageReplyMarkup]: %s\n", err)
		}
	} else {
		timetableName := strings.TrimPrefix(callback_text, "lectures_")
		rows := model.GetLectureYears(callback_text, model.Timetables[timetableName].Course)
		keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)
		editConfig := tgbotapi.NewEditMessageReplyMarkup(chatId, messageId, keyboard)
		_, err := bot.Send(editConfig)
		if err != nil {
			log.Printf("Error [bot.Send() for the NewEditMessageReplyMarkup]: %s\n", err)
		}
	}
}
