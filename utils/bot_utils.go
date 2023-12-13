package utils

import (
	tgbotapi "github.com/samuelemusiani/telegram-bot-api"
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

func GetChatMembers(bot *tgbotapi.BotAPI, chatID int64, memberIds []int64) []tgbotapi.ChatMember {
	var members []tgbotapi.ChatMember

	for _, id := range memberIds {
		chatConfigWithUser := tgbotapi.ChatConfigWithUser{
			ChatID: chatID,
			UserID: id,
		}

		getChatMemberConfig := tgbotapi.GetChatMemberConfig{
			ChatConfigWithUser: chatConfigWithUser,
		}

		member, err := bot.GetChatMember(getChatMemberConfig)
		if err != nil {
			member = makeUnknownMember(chatConfigWithUser)
		}
		members = append(members, member)
	}

	return members
}
