package model

import (
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/text/unicode/norm"
)

// Wrapper for the send function, to send HTML messages
func SendHTML(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	msg.ParseMode = "HTML"
	bot.Send(msg)
}

/* convert a string into kebab case
 * useful for GitHub repository
 *
 * example:
 * string = "Logica per l'informatica"
 * converted_string = ToKebabCase(string); = "logica-per-informatica" (sic!)
 */
func ToKebabCase(str string) string {
	// normalize the string to NFD form
	normalizedStr := norm.NFD.String(strings.ToLower(strings.TrimSpace(str)))

	// remove diacritical marks from the string e.g. à -> a
	reg := regexp.MustCompile(`\p{M}`)
	normalizedStr = reg.ReplaceAllString(normalizedStr, "")

	splitted := strings.Split(normalizedStr, " ")

	// This is not garanteed to work, fix me if error.
	for i := range splitted {
		apostropheSplit := strings.Split(splitted[i], "'")
		splitted[i] = apostropheSplit[len(apostropheSplit)-1]
	}

	return strings.Join(splitted, "-")
}
