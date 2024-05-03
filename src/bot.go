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

	log.Println("ALERT: Connection ended.")
}

func main() {

	log.SetFlags(0)
	godotenv.Load("../config.env")

	client, error := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if error != nil {
		log.Fatalf("Error creating discord session: %v\n", error)
		return
	}

	// --> Maintein the bot online until interrupted
	// --> Close the connection when interrupted
	defer client.Close()
	defer waitUntilInterrupted()

	// --> Handlers
	client.AddHandler(messageReceiver)

	// --> Intents
	client.Identify.Intents = discordgo.IntentGuildMessages
	client.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{Name: "Learning!", Type: discordgo.ActivityTypeGame},
	}

	log.Println("Session created")

	// --> Open the connection, that means the bot will go online
	error = client.Open()
	if error != nil {
		log.Fatalf("Error opening connection: %v\n", error)
		return
	}

	log.Println("Bot is online!")
}

func messageReceiver(s *discordgo.Session, data *discordgo.MessageCreate) {
	if data.Author.ID == s.State.User.ID {
		return
	}
	_, error := s.ChannelMessageSendReply(data.Message.ChannelID, "Hello!", data.Reference())
	if error != nil {
		log.Panic(error)
		return
	}
}
