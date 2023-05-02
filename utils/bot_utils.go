package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func makeUnknownMember(chatConfigWithUser tgbotapi.ChatConfigWithUser) tgbotapi.ChatMember {
	return tgbotapi.ChatMember{
		User: &tgbotapi.User{
			ID:        chatConfigWithUser.UserID,
			FirstName: "???",
			LastName:  "???",
			UserName:  "???",
		},
	}
}

func GetChatMembers(bot *tgbotapi.BotAPI, chatID int64, memberIds []int) ([]tgbotapi.ChatMember, error) {
	var members []tgbotapi.ChatMember

	for _, id := range memberIds {
		chatConfigWithUser := tgbotapi.ChatConfigWithUser{
			ChatID: chatID,
			UserID: id,
		}
		member, err := bot.GetChatMember(chatConfigWithUser)
		if err != nil {
			member = makeUnknownMember(chatConfigWithUser)
		}
		members = append(members, member)
	}

	return members, nil
}
