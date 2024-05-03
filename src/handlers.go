package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func initHandlers() {
	// --> Handlers
	client.AddHandler(messageReceiver)
	client.AddHandler(commandReceiver)
}

func commandReceiver(s *discordgo.Session, data *discordgo.InteractionCreate) {
	s.InteractionRespond(data.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: "Hiii, nice to meet you ^-^ !"},
	})
}

func messageReceiver(s *discordgo.Session, data *discordgo.MessageCreate) {
	if data.Author.ID == s.State.User.ID {
		return
	}
	_, error := s.ChannelMessageSendReply(data.Message.ChannelID, "Hello!", data.Reference())
	if error != nil {
		log.Panic(error)
		return
	}
}
