package commands

import "github.com/bwmarrin/discordgo"

type AstridCommand struct {
	Command *discordgo.ApplicationCommand
	Handler func(*discordgo.Session, *discordgo.InteractionCreate)
}

var ApplicationCommands = []*AstridCommand{
	Hello_Command,
	Hi_Command,
}

var RegisteredCommands = make([]*AstridCommand, len(ApplicationCommands))
