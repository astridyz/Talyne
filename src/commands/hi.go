package commands

import "github.com/bwmarrin/discordgo"

var Hi_Command = &AstridCommand{
	Command: &discordgo.ApplicationCommand{Name: "hi", Description: "Description"},
	Handler: hiMessageReceiver,
}

func hiMessageReceiver(s *discordgo.Session, data *discordgo.InteractionCreate) {
	s.InteractionRespond(data.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: "Hi there! c:"},
	})
}
