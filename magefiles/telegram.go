package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetChat(token string) error {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	updates, err := bot.GetUpdates(tgbotapi.UpdateConfig{})
	if err != nil {
		return err
	}

	for _, update := range updates {
		fmt.Println(update.Message.Chat.ID)
	}

	return nil
}

var ferdz int64 = 638580175
var steph int64 = 6276345756

func SendChat(token string) error {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	_, err = bot.Send(tgbotapi.NewMessage(steph, "I know what you did last night..."))
	if err != nil {
		return err
	}

	return nil
}
