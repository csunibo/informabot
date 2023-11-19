package utils

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"

	tgbotapi "github.com/musianisamuele/telegram-bot-api"
	"golang.org/x/text/unicode/norm"
)

// SendHTML is a wrapper for the send function, to send HTML messages
func SendHTML(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) (tgbotapi.Message, error) {
	msg.ParseMode = tgbotapi.ModeHTML
	msg.DisableWebPagePreview = true
	return bot.Send(msg)
}

/*
ToKebabCase converts a string into kebab case. Useful for GitHub repository
names.
*/
func ToKebabCase(str string) string {
	return toLowerCaseConvention(str, '-')
}

/*
ToSnakeCase converts a string into snake case. Useful for Telegram bots
commands.
*/
func ToSnakeCase(str string) string {
	return toLowerCaseConvention(str, '_')
}

func isSeparator(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsDigit(r)
}

func toLowerCaseConvention(str string, delimiter rune) string {
	// normalize the string to NFD form
	normalizedStr := norm.NFD.String(strings.ToLower(strings.TrimSpace(str)))

	// remove diacritical marks from the string e.g. à -> a
	reg := regexp.MustCompile(`\p{M}`)
	normalizedStr = reg.ReplaceAllString(normalizedStr, "")

	splitted := strings.FieldsFunc(normalizedStr, isSeparator)

	// removing words before "'" character.
	for i := range splitted {
		apostropheSplit := strings.Split(splitted[i], "'")
		splitted[i] = apostropheSplit[len(apostropheSplit)-1]
	}

	return strings.Join(splitted, string(delimiter))
}

func WriteJSONFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	err = file.Close()
	return err
}

const BEGINNING_MONTH = time.September

/*
Returns the starting solar year of the current academic year (e.g. returns 2023
if the current academic year is 2023/24)
*/
func GetCurrentAcademicYear() int {
	now := time.Now()
	year := now.Year()
	if now.Month() >= BEGINNING_MONTH {
		return year
	} else {
		return year - 1
	}
}
