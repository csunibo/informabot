package bot

import (
	"log"
	"strings"

	"github.com/csunibo/informabot/model"
	"github.com/csunibo/informabot/parse"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var autoreplies []model.AutoReply

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
			// commands
			if strings.ToLower(update.Message.Command()) == "start" {
				log.Printf("@%s: \t%s -> %s", update.Message.From.UserName, update.Message.Text, "hello world")
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "hello world")
				bot.Send(msg)
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
