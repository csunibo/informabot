package bot

import (
	"log"
	"strings"

	"github.com/csunibo/informabot/model"
	"github.com/csunibo/informabot/parse"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	autoreplies []model.AutoReply
	actions     []model.Action
)

func StartInformaBot(token string, debug bool) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Error creating new bot: %s", err)
	}
	bot.Debug = debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	autoreplies, err = parse.ParseAutoReplies()
	if err != nil {
		log.Fatalf("Error reading autoreply.json file: %s", err.Error())
	}

	actions, err = parse.ParseActions()
	if err != nil {
		log.Fatalf("Error reading actions.json file: %s", err.Error())
	}

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
			commandName := strings.ToLower(update.Message.Command())

			comparator := func(action model.Action, commandName string) bool {
				return action.Name == commandName
			}
			index := Find(actions, commandName, comparator)
			if index != -1 {
				log.Printf("@%s: \t%s -> %s", update.Message.From.UserName, update.Message.Text, commandName)
				commandName = actions[index].Name
				for commandName != "" {
					commandName = actions[index].Data.HandleBotCommand(bot, update.Message)
					index = Find(actions, commandName, comparator)
				}

				// NOTA: un pattern di questo genere ha senso?
				// invece di chiamare direttamente il metodo su Data, ci teniamo un passaggio di mezzo
				// come se fosse middleware, per cose come log.
				// actions[index].Execute(bot, update)
			} else {
				log.Printf("@%s: \t%s -> COMMAND NOT AVAILABLE", update.Message.From.UserName, update.Message.Text)
			}
		} else {
			// text message
			for i := 0; i < len(autoreplies); i++ {
				if strings.Contains(strings.ToLower(update.Message.Text), strings.ToLower(autoreplies[i].Text)) {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, autoreplies[i].Reply)
					msg.ParseMode = tgbotapi.ModeHTML
					bot.Send(msg)
				}
			}
		}

	}
}

func Find[T any, Q any](a []T, x Q, compare func(T, Q) bool) int {
	for i, n := range a {
		if compare(n, x) {
			return i
		}
	}
	return -1
}
