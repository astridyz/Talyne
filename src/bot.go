package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/astridyz/talyne-discord-bot/commands"

	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var client *discordgo.Session

const TESTING_GUILD_ID string = "1235669274622820362"

func waitUntilInterrupted() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Println("Connection ended.")
}

func removeRegisteredCommands() {
	registeredCommands, error := client.ApplicationCommands(client.State.User.ID, TESTING_GUILD_ID)
	if error != nil {
		log.Panicf("Error getting application commands: %v\n", error)
	}

	for _, ApplicationCommand := range registeredCommands {
		error := client.ApplicationCommandDelete(client.State.User.ID, TESTING_GUILD_ID, ApplicationCommand.ID)
		if error != nil {
			log.Panicf("Error deleting command: %v\n", error)
		}
	}
}

func main() {

	log.SetFlags(0)
	godotenv.Load("../config.env")

	var error error
	client, error = discordgo.New("Bot " + os.Getenv("TOKEN"))
	if error != nil {
		log.Fatalf("Error creating discord session: %v\n", error)
		return
	}

	// --> Maintein the bot online until interrupted
	// --> Delete all registered commands *testing function*
	// --> Close the connection when interrupted
	defer client.Close()
	defer removeRegisteredCommands()
	defer waitUntilInterrupted()

	// --> Intents
	client.Identify.Intents = discordgo.IntentGuildMessages
	client.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{Name: "Learning!", Type: discordgo.ActivityTypeGame},
	}

	log.Println("Session created")
	// --> Starting all events and commands
	initHandlers()
	commands.Init()

	// --> Open the connection, that means the bot will go online
	error = client.Open()
	if error != nil {
		log.Fatalf("Error opening connection: %v\n", error)
		return
	}

	// --> Creating all slash commands
	for i, AstridCommand := range commands.ApplicationCommands {
		_, error = client.ApplicationCommandCreate(client.State.User.ID, TESTING_GUILD_ID, AstridCommand.Command)
		if error != nil {
			log.Panicf("Error creating bot command: %v\n", error)
			continue
		}

		commands.RegisteredCommands[i] = AstridCommand
	}

	log.Println("Bot is online!")
}
