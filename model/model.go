package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/csunibo/informabot/commands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type DataInterface interface {
	// Returns another command to be executed, or emtpy string if no command is to be executed
	// NOTE: we assume that everything returned by this function is a valid command
	HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string
}

func GetActionFromType(name string, commandType string) Action {
	var data DataInterface
	switch commandType {
	case "message":
		data = MessageData{}
	case "help":
		data = HelpData{}
	case "update":
		data = UpdateData{}
	case "lookingFor":
		data = LookingForData{}
	case "notLookingFor":
		data = NotLookingForData{}
	case "yearly":
		data = YearlyData{}
	case "todayLectures":
		data = TodayLecturesData{}
	case "tomorrowLectures":
		data = TomorrowLecturesData{}
	case "list":
		data = ListData{}
	case "course":
		data = CourseData{}
	case "luck":
		data = LuckData{}
	default:
		data = InvalidData{}
	}

	return Action{
		Name: name,
		Type: commandType,
		Data: data,
	}
}

type AutoReply struct {
	Text  string `json:"text"`
	Reply string `json:"reply"`
}

type Action struct {
	Name string
	Type string        `json:"type"`
	Data DataInterface `json:"data"`
}

type MessageData struct {
	Text string `json:"text"`
}

type HelpData struct {
	Description string `json:"description"`
}

type UpdateData struct {
	Description string `json:"description"`
	NoYear      string `json:"noYear"`
	NoMod       string `json:"noMod"`
	Started     string `json:"started"`
	Ended       string `json:"ended"`
	Failed      string `json:"failed"`
}

type LookingForData struct {
	Description  string `json:"description"`
	SingularText string `json:"singularText"`
	PluralText   string `json:"pluralText"`
	ChatError    string `json:"chatError"`
}

type NotLookingForData struct {
	Description   string `json:"description"`
	Text          string `json:"text"`
	ChatError     string `json:"chatError"`
	NotFoundError string `json:"notFoundError"`
}

type YearlyData struct {
	Description string `json:"description"`
	Command     string `json:"command"`
	NoYear      string `json:"noYear"`
}

type TodayLecturesData struct {
	Description  string `json:"description"`
	Url          string `json:"url"`
	Title        string `json:"title"`
	FallbackText string `json:"fallbackText"`
}

type TomorrowLecturesData TodayLecturesData

type ListData struct {
	Description string     `json:"description"`
	Header      string     `json:"header"`
	Template    string     `json:"template"`
	Items       [][]string `json:"items"`
}

type CourseData struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Virtuale    string   `json:"virtuale"`
	Teams       string   `json:"teams"`
	Website     string   `json:"website"`
	Professors  []string `json:"professors"`
	Telegram    string   `json:"telegram"`
}

type LuckData struct {
	Description string `json:"description"`
}

type InvalidData struct{}

func SendHTML(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	msg.ParseMode = "HTML"
	bot.Send(msg)
}

func (data MessageData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, data.Text)
	SendHTML(bot, msg)

	return ""
}

func (data HelpData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO HelpData: notimplemented, Got: %s\n", message.Text))
	SendHTML(bot, msg)

	return ""
}

func (data UpdateData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO UpdateData: notimplemented, Got: %s\n", message.Text))
	SendHTML(bot, msg)

	return ""
}

func (data LookingForData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO LookingForData: notimplemented, Got: %s\n", message.Text))
	SendHTML(bot, msg)

	return ""
}

func (data NotLookingForData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO NotLookingForData: notimplemented, Got: %s\n", message.Text))
	SendHTML(bot, msg)

	return ""
}

func (data YearlyData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	var chatTitle string = message.Chat.Title

	// check if string starts with "Yearly"
	if strings.HasPrefix(chatTitle, "primo") {
		return data.Command + "1"
	} else if strings.HasPrefix(chatTitle, "secondo") {
		return data.Command + "2"
	} else if strings.HasPrefix(chatTitle, "terzo") {
		return data.Command + "3"
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, data.NoYear)
		SendHTML(bot, msg)
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
	SendHTML(bot, msg)

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
	SendHTML(bot, msg)

	return ""
}

func (data ListData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO ListData: notimplemented, Got: %s\n", message.Text))
	SendHTML(bot, msg)

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
			ternary_assignment(data.Name != "", fmt.Sprintf("<a href='https://csunibo.github.io/%s/'>📚 Risorse: materiali, libri, prove</a>\n", ToKebabCase(data.Name)))+
			ternary_assignment(data.Name != "", fmt.Sprintf("<a href='https://github.com/csunibo/%s/'>📂 Repository GitHub delle risorse</a>\n", ToKebabCase(data.Name)))+
			ternary_assignment(data.Telegram != "", fmt.Sprintf("<a href='t.me/$%s'>👥 Gruppo Studenti</a>\n", data.Telegram)))
	SendHTML(bot, msg)

	return ""
}

func (data LuckData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO LuckData: notimplemented, Got: %s\n", message.Text))
	SendHTML(bot, msg)

	return ""
}

func (data InvalidData) HandleBotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) string {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TODO InvalidData: notimplemented, Got: %s\n", message.Text))
	SendHTML(bot, msg)

	return ""
}

// TODO: sposta questa funzione nella sua libreria
/* convert a string into kebab case
 * useful for GitHub repository
 *
 * example:
 * string = "Logica per l'informatica"
 * converted_string = ToKebabCase(string); = "logica-per-informatica" (sic!)
 */
func ToKebabCase(str string) string {
	// normalize the string to NFD form
	norm := strings.ToLower(strings.TrimSpace(str))

	splitted := strings.Split(norm, " ")

	// This is not garanteed to work, fix me if error.
	for i := range splitted {
		apostropheSplit := strings.Split(splitted[i], "'")
		splitted[i] = apostropheSplit[len(apostropheSplit)-1]
	}

	return strings.Join(splitted, "-")
}
