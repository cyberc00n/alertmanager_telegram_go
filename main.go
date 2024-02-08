package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Alert struct {
	Status      string            `json:"status"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type WebhookMessage struct {
	Alerts []Alert `json:"alerts"`
}

var (
	telegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID           = os.Getenv("TELEGRAM_CHAT_ID")
)

func main() {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Bot successfully started")

	http.HandleFunc("/alert", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received new request")

		var m WebhookMessage
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			log.Printf("Could not decode message: %v", err)
			http.Error(w, "Could not decode message", http.StatusBadRequest)
			return
		}

		log.Printf("Message decoded successfully")

		for _, alert := range m.Alerts {
			msg := formatAlertMessage(alert)
			sendMessage(bot, msg)
		}

		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func formatAlertMessage(alert Alert) string {
	statusMessage := ""
	switch alert.Status {
	case "resolved":
		statusMessage = "Service recovered: "
	case "firing":
		statusMessage = "Alert triggered: "
	default:
		statusMessage = "Alert received: "
	}
	return statusMessage + alert.Annotations["description"]
}

func sendMessage(bot *tgbotapi.BotAPI, msg string) {
	log.Printf("Sending message: %s", msg)
	message := tgbotapi.NewMessageToChannel(chatID, msg)
	if _, err := bot.Send(message); err != nil {
		log.Printf("Could not send message: %v", err)
	}
}
