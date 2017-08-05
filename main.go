package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/novikk/splitbot/hutoma"
	"github.com/novikk/splitbot/webhooks"

	"gopkg.in/telegram-bot-api.v4"
)

var HUTOMA_BOT_ID string
var HUTOMA_DEV_KEY string
var HUTOMA_CLIENT_KEY string

var TELEGRAM_TOKEN string

var lastSpeaker string

func init() {
	// init hutoma vars from environment
	HUTOMA_BOT_ID = os.Getenv("HUTOMA_BOT_ID")
	HUTOMA_DEV_KEY = os.Getenv("HUTOMA_DEV_KEY")
	HUTOMA_CLIENT_KEY = os.Getenv("HUTOMA_CLIENT_KEY")

	// init telegram
	TELEGRAM_TOKEN = os.Getenv("TELEGRAM_TOKEN")
}

func startTelegramBot() {
	hc := hutoma.HutomaClient{HUTOMA_BOT_ID, HUTOMA_DEV_KEY, HUTOMA_CLIENT_KEY, ""}

	bot, err := tgbotapi.NewBotAPI(TELEGRAM_TOKEN)
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
		} else {
			fmt.Println("?", update.Message)
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// save last speaker
		lastSpeaker = update.Message.From.FirstName + " " + update.Message.From.LastName
		if lastSpeaker == "" {
			lastSpeaker = update.Message.From.UserName
		}

		if lastSpeaker == "" {
			lastSpeaker = strconv.Itoa(update.Message.From.ID)
		}

		webhooks.SetLastSpeaker(lastSpeaker)

		// send the message to hutoma
		hres, err := hc.Chat(update.Message.Text)
		if err != nil {
			log.Printf("Error chatting: %s\n", err)
			continue
		}

		log.Printf("[Hutoma] %s", hres.Result.Answer)
		if hres.Result.Answer == "unknown" {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, hres.Result.Answer)
		// msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func startWebhooks() {
	webhooksRouter := mux.NewRouter().StrictSlash(true)
	webhooksRouter.HandleFunc("/expense", webhooks.HandleExpense)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), webhooksRouter))
}

func main() {
	go startTelegramBot()
	startWebhooks()
}
