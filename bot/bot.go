package bot

import (
	"log"
	"strings"

	"github.com/csunibo/informabot/model"
	"github.com/csunibo/informabot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Error getting updates: %s", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			handleCommand(bot, &update)
		} else {
			// text message
			for i := 0; i < len(model.Autoreplies); i++ {
				if strings.Contains(strings.ToLower(update.Message.Text), strings.ToLower(model.Autoreplies[i].Text)) {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, model.Autoreplies[i].Reply)
					msg.ParseMode = tgbotapi.ModeHTML
					bot.Send(msg)
				}
			}
		}

	}
}

func handleCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	commandName := strings.ToLower(update.Message.Command())

	hasExecutedCommand := executeCommandWithName(bot, update, commandName)
	if !hasExecutedCommand {
		memeIndex := utils.Find(model.MemeList, commandName, func(meme model.Meme, commandName string) bool {
			return meme.Name == commandName
		})

		if memeIndex != -1 {
			log.Printf("@%s: \t%s -> MEMES", update.Message.From.UserName, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, model.MemeList[memeIndex].Text)
			bot.Send(msg)
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
		if newCommand != "" {
			// NOTA: un pattern di questo genere ha senso?
			// invece di chiamare direttamente il metodo su Data, ci teniamo un passaggio di mezzo
			// come se fosse middleware, per cose come log.
			// actions[index].Execute(bot, update)
			executeCommandWithName(bot, update, newCommand)
		}
	}
}

// executes a given command in the command list, given its name
// @return true if command was found, false otherwise
func executeCommandWithName(bot *tgbotapi.BotAPI, update *tgbotapi.Update, commandName string) bool {
	idx := utils.Find(model.Actions, commandName, func(action model.Action, commandName string) bool {
		return action.Name == commandName
	})

	if idx != -1 {
		executeCommand(bot, update, idx)
		return true
	}

	return false
}
