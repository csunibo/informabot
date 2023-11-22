package model

import (
	"fmt"
	"strings"

	tgbotapi "github.com/musianisamuele/telegram-bot-api"
)

type InlineKeyboardRows [][]tgbotapi.InlineKeyboardButton

// Returns a group of button rows for a selected groups on `timetables`
func GetTimetableCoursesRows(timetables *map[string]Timetable) InlineKeyboardRows {
	rows := make([][]tgbotapi.InlineKeyboardButton, len(*timetables))

	i := 0
	for callback, timetable := range *timetables {
		row := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(timetable.Course, fmt.Sprintf("lectures_%s", callback)))
		rows[i] = row
		i++
	}

	return rows
}

// Returns buttons which permits to choose the day for the timetable
func ChooseTimetableDay(callback_text string) InlineKeyboardRows {
	rows := make([][]tgbotapi.InlineKeyboardButton, 2)
	rows[0] = tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Oggi", fmt.Sprintf("%s_today", callback_text)))
	rows[1] = tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Domani", fmt.Sprintf("%s_tomorrow", callback_text)))

	return rows
}

// Returns a group of buttons rows for the available years of a `course`
func GetLectureYears(callback_text string, course string) InlineKeyboardRows {
	yearsNro := 3
	// Master degrees has a duration of only 2 years
	if strings.HasPrefix(callback_text, "lectures_lm") {
		yearsNro = 2
	}
	rows := make([][]tgbotapi.InlineKeyboardButton, yearsNro)

	i := 1
	for i <= yearsNro {
		buttonText := fmt.Sprintf("%s: %d^ anno", course, i)
		buttonCallback := fmt.Sprintf("%s_y_%d", callback_text, i)
		row := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(buttonText, buttonCallback))
		rows[i-1] = row

		i++
	}

	return rows
}
