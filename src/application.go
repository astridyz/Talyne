package main

import (
	"log"

	"github.com/astridyz/talyne-discord-bot/commands"
)

func removeRegisteredCommands() {
	ApplicationCommands, error := client.ApplicationCommands(client.State.User.ID, TESTING_GUILD_ID)
	if error != nil {
		log.Panicf("Error getting application commands: %v\n", error)
	}

	for _, ApplicationCommand := range ApplicationCommands {
		error := client.ApplicationCommandDelete(client.State.User.ID, TESTING_GUILD_ID, ApplicationCommand.ID)
		if error != nil {
			log.Panicf("Error deleting command: %v\n", error)
		}
	}

	log.Println("Commands deleted")
}

func createApplicationCommands() {
	for i, AstridCommand := range commands.ApplicationCommands {
		_, error := client.ApplicationCommandCreate(client.State.User.ID, TESTING_GUILD_ID, AstridCommand.Command)
		if error != nil {
			log.Panicf("Error creating bot command: %v\n", error)
			continue
		}

		commands.RegisteredCommands[i] = AstridCommand
	}

	log.Println("Commands added")
}
