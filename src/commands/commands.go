package commands

import "github.com/bwmarrin/discordgo"

type Command struct {
	Command *discordgo.ApplicationCommand
	Handler func(*discordgo.Session, *discordgo.InteractionCreate)
}

var ApplicationCommands = []*Command{
	Hello_Command,
	Hi_Command,
}

var RegisteredCommands = make([]*Command, len(ApplicationCommands))
