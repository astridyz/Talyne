package commands

import "github.com/bwmarrin/discordgo"

type AstridCommand struct {
	Command *discordgo.ApplicationCommand
	Handler func(*discordgo.Session, *discordgo.InteractionCreate)
}

var applicationCommands []*AstridCommand
var RegisteredCommands = make([]*AstridCommand, len(applicationCommands))

func GetAllCommands() []*AstridCommand {
	applicationCommands = []*AstridCommand{
		Hello_Command,
		Hi_Command,
	}

	return applicationCommands
}
