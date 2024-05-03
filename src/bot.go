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

func waitUntilInterrupted() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Println("Connection ended.")
}

func remoteApplicationsCommands() {
	registeredCommands, error := client.ApplicationCommands(client.State.User.ID, "1235669274622820362")
	if error != nil {
		log.Fatalf("Could not fetch registered commands: %v", error)
	}

	for _, command := range registeredCommands {
		error := client.ApplicationCommandDelete(client.State.User.ID, "1235669274622820362", command.ID)
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
	defer remoteApplicationsCommands()
	defer waitUntilInterrupted()

	// --> Intents
	client.Identify.Intents = discordgo.IntentGuildMessages
	client.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{Name: "Learning!", Type: discordgo.ActivityTypeGame},
	}

	log.Println("Session created")
	// --> Starting all events
	initHandlers()

	// --> Open the connection, that means the bot will go online
	error = client.Open()
	if error != nil {
		log.Fatalf("Error opening connection: %v\n", error)
		return
	}

	// --> Creating a command
	//_, error = client.ApplicationCommandCreate(client.State.User.ID, "1235669274622820362", commands.Hello_Command)
	//if error != nil {
	//	log.Panicf("Error creating bot command: %v\n", error)
	//}

	for _, AstridCommand := range commands.GetAllCommands() {
		_, error = client.ApplicationCommandCreate(client.State.User.ID, "1235669274622820362", AstridCommand.Command)
		if error != nil {
			log.Panicf("Error creating bot command: %v\n", error)
		}
	}

	log.Println("Bot is online!")
}
