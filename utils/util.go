package utils

import (
	"encoding/json"
	"os"
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

// TODO: this function is already implemented in Slices Index, use that.
func Find[T any, Q any](a []T, x Q, compare func(T, Q) bool) int {
	for i, n := range a {
		if compare(n, x) {
			return i
		}
	}
	return -1
}

func WriteJSONFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
