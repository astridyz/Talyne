package commands

import "github.com/bwmarrin/discordgo"

type AstridCommand struct {
	Command *discordgo.ApplicationCommand
	Handler func(*discordgo.Session, *discordgo.InteractionCreate)
}

var ApplicationCommands []*AstridCommand
var RegisteredCommands []*AstridCommand

func Init() {
	ApplicationCommands = []*AstridCommand{
		Hello_Command,
		Hi_Command,
	}

	RegisteredCommands = make([]*AstridCommand, len(ApplicationCommands))
}
