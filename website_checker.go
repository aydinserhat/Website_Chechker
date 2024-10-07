package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	botToken := "Your bot token"
	id := int64(Your id)

	bot, hata := tgbotapi.NewBotAPI(botToken)
	if hata != nil {
		log.Panic(hata)
	}
	for {
		url := "https://x.com/home"

		yanitSuresi := time.Now()
		client := http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.Get(url)
		if err != nil {
			message := fmt.Sprintf("Site aktif değil: %v", err)
			sendTelegramMessage(bot, id, message)
			return
		}
		defer resp.Body.Close()
		istek_suresi := time.Since(yanitSuresi)
		var message string
		if resp.StatusCode == http.StatusOK {
			istek_sure := istek_suresi.Milliseconds()
			if istek_sure > 100 {
				message = fmt.Sprintf("up %v ms- Yavaş", istek_sure)
			} else {
				message = fmt.Sprintf("up %v ms", istek_sure)
			}

		} else {
			message = fmt.Sprintf("down - %d", resp.StatusCode)
		}
		sendTelegramMessage(bot, id, message)
		time.Sleep(5 * time.Minute)
	}
}

func sendTelegramMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
	mesaj := tgbotapi.NewMessage(chatID, message)
	_, err := bot.Send(mesaj)
	if err != nil {
		fmt.Println("Mesaj gönderilemedi:", err)
	} else {
		fmt.Println("Telegram'a gönderildi")
	}

}
