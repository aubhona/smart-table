package app

import (
	"strconv"

	"gopkg.in/telebot.v4"
)

type BotQueryService struct {
	bot *telebot.Bot
}

func NewBotQueryService(bot *telebot.Bot) *BotQueryService {
	return &BotQueryService{bot: bot}
}

func (q *BotQueryService) SendMessage(chatID, message string) error {
	chaIDInt, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		return err
	}

	_, err = q.bot.Send(&telebot.User{ID: chaIDInt}, message)

	return err
}
