package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/astridyz/talyne-discord-bot/handlers"

	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var client *discordgo.Session

const TESTING_GUILD_ID = "1235669274622820362"

func main() {

	log.SetFlags(0)
	godotenv.Load("../config.env")

	var error error
	client, error = discordgo.New("Bot " + os.Getenv("TOKEN"))
	if error != nil {
		log.Fatalf("Error creating discord session: %v\n", error)
		return
	}

	log.Println("Session created")

	// --> Closing the bot
	defer client.Close()
	defer removeRegisteredCommands()
	defer waitUntilInterrupted()

	// --> Intents
	client.Identify.Intents = discordgo.IntentGuildMessages
	client.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{Name: "Learning!", Type: discordgo.ActivityTypeGame},
	}

	// --> Starting handlers
	client.AddHandler(handlers.CommandHandler)

	// --> Open the connection, that means the bot will go online
	error = client.Open()
	if error != nil {
		log.Fatalf("Error opening connection: %v\n", error)
		return
	}

	// --> Creating all slash commands
	createApplicationCommands()

	log.Println("Bot is online!")
}

func waitUntilInterrupted() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Println("Connection ended.")
}
