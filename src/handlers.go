package main

import (
	"log"

	"github.com/astridyz/talyne-discord-bot/commands"
	"github.com/bwmarrin/discordgo"
)

func initHandlers() {
	// --> Handlers
	client.AddHandler(messageReceiver)
	client.AddHandler(commandReceiver)
}

func commandReceiver(s *discordgo.Session, data *discordgo.InteractionCreate) {
	for _, AstridCommand := range commands.RegisteredCommands {
	}
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
