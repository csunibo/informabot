package utils

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"

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
ToKebabCase convert a string into kebab case. Useful for GitHub repository
names.
*/
func ToKebabCase(str string) string {
	// normalize the string to NFD form
	normalizedStr := norm.NFD.String(strings.ToLower(strings.TrimSpace(str)))

	// remove diacritical marks from the string e.g. à -> a
	reg := regexp.MustCompile(`\p{M}`)
	normalizedStr = reg.ReplaceAllString(normalizedStr, "")

	splitted := strings.Split(normalizedStr, " ")

	// removing words before "'" character.
	for i := range splitted {
		apostropheSplit := strings.Split(splitted[i], "'")
		splitted[i] = apostropheSplit[len(apostropheSplit)-1]
	}

	return strings.Join(splitted, "-")
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
