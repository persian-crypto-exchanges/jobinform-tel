package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gravelstone/gravel"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	t := os.Getenv("TELEGRAM_TOKEN")
	if t == "" {
		log.Fatal("TELEGRAM_TOKEN is not set")
	}

	client := gravel.NewGravel(t, true)

	for {
		updates, err := client.GetUpdates()
		if err != nil {
			log.Printf("Error fetching updates: %v", err)
			continue
		}

		for _, update := range updates {
			if update.Message != nil {
				if update.Message.IsCommand() {
					err = client.SendMessage(update.Message.Chat.ID, "Hello! You sent: "+update.Message.Text)
				}

				if update.Message.Text == "ping" {
					usr, _ := client.GetUserInfo(update.Message.Chat.ID)
					client.SendMessage(update.Message.Chat.ID, fmt.Sprintf("Pong! Hello, %v!", usr.FirstName))
				}

				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
			}
		}
	}

}
