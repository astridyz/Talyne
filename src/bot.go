package main

import (
	"os"
	"os/signal"
	"syscall"

	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func waitUntilInterrupted() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func main() {

	log.SetFlags(0)
	godotenv.Load("../config.env")

	client, error := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if error != nil {
		log.Fatalf("Error creating discord session: %v\n", error)
		return
	}

	client.Identify.Intents = discordgo.IntentGuildMessages
	client.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{Name: "Learning!", Type: discordgo.ActivityTypeGame},
	}

	defer waitUntilInterrupted()

	log.Println("Session created")

	error = client.Open()
	if error != nil {
		log.Fatalf("Error opening connection: %v\n", error)
		return
	}

	log.Println("Bot is online!")
}
