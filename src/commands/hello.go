package commands

import "github.com/bwmarrin/discordgo"

var Hello_Command = &AstridCommand{
	Command: &discordgo.ApplicationCommand{Name: "hello", Description: "Description"},
	Handler: helloMessageReceiver,
}

func helloMessageReceiver(s *discordgo.Session, data *discordgo.InteractionCreate) {
	s.InteractionRespond(data.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: "Hello, nice to meet you ^-^ !"},
	})
}
