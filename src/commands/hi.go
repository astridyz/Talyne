package commands

import (
	"github.com/astridyz/talyne-discord-bot/utils"
	"github.com/bwmarrin/discordgo"
)

var Hi_Command = &AstridCommand{
	Command: &discordgo.ApplicationCommand{Name: "hi", Description: "Description"},
	Handler: hiMessageReceiver,
}

func hiMessageReceiver(s *discordgo.Session, data *discordgo.InteractionCreate) {
	var Interaction = utils.AstridInteraction{Client: s, Data: data}
	Interaction.SendMessage("Hii!! Are you okay?")
}
