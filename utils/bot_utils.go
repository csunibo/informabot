package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func GetChatMembers(bot *tgbotapi.BotAPI, chatID int64, memberIds []int) ([]tgbotapi.ChatMember, error) {
	var members []tgbotapi.ChatMember

	for _, id := range memberIds {
		member, err := bot.GetChatMember(tgbotapi.ChatConfigWithUser{
			ChatID: chatID,
			UserID: id,
		})
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}
