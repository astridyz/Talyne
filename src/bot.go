package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/astridyz/talyne-discord-bot/handlers"

	aura "github.com/astridyz/Aura/src"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var client *discordgo.Session

var log = aura.New(&aura.Prefix{
	Body:  "Astrid:",
	Level: aura.Debug,
})

const TESTING_GUILD_ID = "1235669274622820362"

func main() {

	godotenv.Load("../config.env")

	var error error
	client, error = discordgo.New("Bot " + os.Getenv("TOKEN"))
	if error != nil {
		log.Fatalf("Error creating discord session: %v\n", error)
		return
	}

	// --> Closing the bot
	defer client.Close()
	defer removeRegisteredCommands()
	defer waitUntilInterrupted()

	// --> Intents
	client.Identify.Intents = discordgo.IntentGuildMessages | discordgo.IntentGuildMembers
	client.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{Name: "Learning!", Type: discordgo.ActivityTypeGame},
	}

	// --> Starting handlers
	client.AddHandler(handlers.CommandHandler)
	client.AddHandler(handlers.WelcomeHandler)

	// --> Open the connection, that means the bot will go online
	error = client.Open()
	if error != nil {
		log.Fatalf("Error opening connection: %v\n", error)
		return
	}

	// --> Creating all slash commands
	createApplicationCommands()

	log.Print("Bot is online!")
}

func waitUntilInterrupted() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Warn("Connection ended.")
}
