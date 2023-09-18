package utils

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"

	tgbotapi "github.com/musianisamuele/telegram-bot-api"
	"golang.org/x/text/unicode/norm"
)

// Wrapper for the send function, to send HTML messages
func SendHTML(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	msg.ParseMode = tgbotapi.ModeHTML
	msg.DisableWebPagePreview = true
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

		/*
			If a "'" character is present, it should be just one character (so two
			words), then there are two cases:
			1) the first word is a single-character word -> the word is discarded
			2) the first word is not a single-character word -> the last character
			of the word is replaced by a "-" character
		*/
		if len(apostropheSplit) == 2 {
			last := len(apostropheSplit[0]) - 1

			// Discarding the last character
			remaining := apostropheSplit[0][:last]
			if remaining == "" {
				// Single-character word case, it is discarded
				splitted[i] = apostropheSplit[1]
			} else {
				// Multiple character word case, adding "-" character to the word
				splitted[i] = remaining + "-" + apostropheSplit[1]
			}
		}
	}

	return strings.Join(splitted, "-")
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
