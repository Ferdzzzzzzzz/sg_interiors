package alert

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Alerter interface {
	Contact(name, email, org, message string) error
}

type LogOnly struct{}

func (t LogOnly) Contact(name, email, org, message string) error {
	fmt.Printf(`
name: %s, 
email: %s, 
org: %s, 
message: %s

`, name, email, org, message)

	return nil
}

type Bot struct {
	bot *tgbotapi.BotAPI
	l   *log.Logger
}

func NewTelegramAlertBot(
	telegramApiToken string,
	l *log.Logger,
) (Bot, error) {
	bot, err := tgbotapi.NewBotAPI(telegramApiToken)
	if err != nil {
		return Bot{}, err
	}
	return Bot{bot, l}, nil
}

var ferdz int64 = 638580175
var stevo int64 = 933906334

var numbers = []int64{ferdz, stevo}

func (a Bot) Contact(name, email, org, message string) error {
	a.l.Println("sending alert to Telegram alert group")

	msg := fmt.Sprintf(`name: %s, 
email: %s, 
org: %s, 
message: %s`, name, email, org, message)

	var maybeErr error
	for _, n := range numbers {
		if _, err := a.bot.Send(tgbotapi.NewMessage(int64(n), msg)); err != nil {
			err = maybeErr
		}
	}

	return maybeErr
}
