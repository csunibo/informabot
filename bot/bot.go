package bot

import (
	"log"
	"strings"

	"github.com/csunibo/informabot/model"
	"github.com/csunibo/informabot/utils"
	tgbotapi "github.com/musianisamuele/telegram-bot-api"
	"golang.org/x/exp/slices"
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

func handleCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	commandName := strings.ToLower(update.Message.Command())

	hasExecutedCommand := executeCommandWithName(bot, update, commandName)
	if !hasExecutedCommand {
		memeIndex := slices.IndexFunc(model.MemeList, func(meme model.Meme) bool {
			return strings.ToLower(meme.Name) == commandName
		})

		if memeIndex != -1 {
			log.Printf("@%s: \t%s -> MEMES", update.Message.From.UserName, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, model.MemeList[memeIndex].Text)
			utils.SendHTML(bot, msg)
		} else {
			executeCommandWithName(bot, update, "unknown")
			log.Printf("@%s: \t%s -> COMMAND NOT AVAILABLE", update.Message.From.UserName, update.Message.Text)
		}
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
			executeCommandWithName(bot, update, newCommand.NextCommand)
		}
	}
}

// executes a given command in the command list, given its name
// @return true if command was found, false otherwise
func executeCommandWithName(bot *tgbotapi.BotAPI, update *tgbotapi.Update, commandName string) bool {
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
