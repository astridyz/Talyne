package commands

import "github.com/bwmarrin/discordgo"

type AstridCommand struct {
	Command *discordgo.ApplicationCommand
	Handler func(*discordgo.Session, *discordgo.InteractionCreate)
}

func GetAllCommands() []*AstridCommand {
	applicationCommands := []*AstridCommand{
		Hello_Command,
		Hi_Command,
	}

	return applicationCommands
}
