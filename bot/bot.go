package bot

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/musianisamuele/telegram-bot-api"
	"golang.org/x/exp/slices"

	"github.com/csunibo/informabot/model"
	"github.com/csunibo/informabot/utils"
)

func StartInformaBot(token string, debug bool) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Error creating new bot: %s", err)
	}
	bot.Debug = debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	model.InitGlobals()

	run(bot)
}

func run(bot *tgbotapi.BotAPI) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		} else if filterMessage(bot, update.Message) {
			continue
		}

		if update.Message.IsCommand() {
			handleCommand(bot, &update)
		} else {
			// text message
			for i := 0; i < len(model.Autoreplies); i++ {
				if strings.Contains(strings.ToLower(update.Message.Text),
					strings.ToLower(model.Autoreplies[i].Text)) {
					var msg tgbotapi.MessageConfig

					if update.Message.IsTopicMessage {
						msg = tgbotapi.NewThreadMessage(update.Message.Chat.ID,
							update.Message.MessageThreadID, model.Autoreplies[i].Reply)
					} else {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID,
							model.Autoreplies[i].Reply)
					}
					msg.ReplyToMessageID = update.Message.MessageID
					utils.SendHTML(bot, msg)
				}
			}
		}

	}
}

type handler = func(*tgbotapi.BotAPI, *tgbotapi.Update, string) bool

var handlers = []handler{handleAction, handleTeaching, handleMeme}

func handleCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	commandName := strings.ToLower(update.Message.Command())
	for _, h := range handlers {
		if h(bot, update, commandName) {
			return
		}
	}
}

func handleTeaching(bot *tgbotapi.BotAPI, update *tgbotapi.Update, teachingName string) bool {
	teaching, ok := model.Teachings[teachingName]
	if !ok {
		return false
	}
	currentAcademicYear := fmt.Sprint(utils.GetCurrentAcademicYear())
	var b strings.Builder
	if teaching.Name != "" {
		b.WriteString(fmt.Sprintf("<b>%s</b>\n", teaching.Name))
	}
	if teaching.Website != "" {
		b.WriteString(fmt.Sprintf("<a href='https://www.unibo.it/it/didattica/insegnamenti/insegnamento/%s/%s'>Sito</a>\n",
			currentAcademicYear, teaching.Website))
		b.WriteString(fmt.Sprintf("<a href='https://www.unibo.it/it/didattica/insegnamenti/insegnamento/%s/%s/orariolezioni'>Orario</a>\n",
			currentAcademicYear, teaching.Website))
	}
	if teaching.Professors != nil {
		emails := strings.Join(teaching.Professors, "@unibo.it\n ") + "@unibo.it\n"
		b.WriteString(fmt.Sprintf("Professori:\n %s", emails))
	}

	if teaching.Name != "" {
		b.WriteString(fmt.Sprintf("<a href='https://risorse.students.cs.unibo.it/%s/'>📚 Risorse (istanza principale)</a>\n", teaching.Url))
		b.WriteString(fmt.Sprintf("<a href='https://dynamik.vercel.app/%s/'>📚 Risorse (istanza di riserva 1)</a>\n", teaching.Url))
		b.WriteString(fmt.Sprintf("<a href='https://csunibo.github.io/dynamik/%s/'>📚 Risorse (istanza di riserva 2)</a>\n", teaching.Url))
		b.WriteString(fmt.Sprintf("<a href='https://github.com/csunibo/%s/'>📂 Repository GitHub delle risorse</a>\n", teaching.Url))
	}
	if teaching.Chat != "" {
		b.WriteString(fmt.Sprintf("<a href='https://t.me/%s'>👥 Gruppo Studenti</a>\n", teaching.Chat))
	}
	utils.SendHTML(bot, tgbotapi.NewMessage(update.Message.Chat.ID, b.String()))
	return true
}

func handleMeme(bot *tgbotapi.BotAPI, update *tgbotapi.Update, memeName string) bool {
	memeIndex := slices.IndexFunc(model.MemeList, func(meme model.Meme) bool {
		return strings.ToLower(meme.Name) == memeName
	})

	if memeIndex != -1 {
		log.Printf("@%s: \t%s -> MEMES", update.Message.From.UserName, update.Message.Text)

		var msg tgbotapi.MessageConfig
		if update.Message.IsTopicMessage {
			msg = tgbotapi.NewThreadMessage(update.Message.Chat.ID,
				update.Message.MessageThreadID, model.MemeList[memeIndex].Text)
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, model.MemeList[memeIndex].Text)
		}
		utils.SendHTML(bot, msg)
		return true
	} else {
		handleAction(bot, update, "unknown")
		log.Printf("@%s: \t%s -> COMMAND NOT AVAILABLE", update.Message.From.UserName, update.Message.Text)
		return false
	}
}

// executes a given command in the command list, given its index
// if invalid index, does nothing
func executeCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update, commandIndex int) {
	if commandIndex >= 0 && commandIndex < len(model.Actions) {
		log.Printf("@%s: \t%s -> COMMAND", update.Message.From.UserName, update.Message.Text)
		newCommand := model.Actions[commandIndex].Data.HandleBotCommand(bot, update.Message)

		if newCommand.HasText() {
			var msg tgbotapi.MessageConfig

			if update.Message.IsTopicMessage {
				msg = tgbotapi.NewThreadMessage(update.Message.Chat.ID,
					update.Message.MessageThreadID, newCommand.Text)
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, newCommand.Text)
			}
			utils.SendHTML(bot, msg)
		}

		if newCommand.HasNextCommand() {
			handleAction(bot, update, newCommand.NextCommand)
		}
	}
}

// executes a given command in the command list, given its name
// @return true if command was found, false otherwise
func handleAction(bot *tgbotapi.BotAPI, update *tgbotapi.Update, commandName string) bool {
	idx := slices.IndexFunc(model.Actions, func(action model.Action) bool {
		return action.Name == commandName
	})

	if idx != -1 {
		executeCommand(bot, update, idx)
		return true
	}

	return false
}

func filterMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) bool {
	if message.Dice != nil {
		// msg := tgbotapi.NewMessage(message.Chat.ID, "Found a dice")
		// bot.Send(msg)
		return true
	}
	return false
}
