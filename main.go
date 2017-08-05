package main

import (
	"log"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

var HUTOMA_BOT_ID string
var HUTOMA_DEV_KEY string
var HUTOMA_CLIENT_KEY string

func init() {
	// init hutoma vars from environment
	HUTOMA_BOT_ID = os.Getenv("HUTOMA_BOT_ID")
	HUTOMA_DEV_KEY = os.Getenv("HUTOMA_DEV_KEY")
	HUTOMA_CLIENT_KEY = os.Getenv("HUTOMA_CLIENT_KEY")
}

func main() {
	bot, err := tgbotapi.NewBotAPI("426886257:AAFevPoGCZlgDVPfeIrmjLLn-hRZyBw8boU")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
