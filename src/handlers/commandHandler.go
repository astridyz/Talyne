package handlers

import (
	"github.com/astridyz/talyne-discord-bot/commands"
	"github.com/bwmarrin/discordgo"
)

func CommandHandler(s *discordgo.Session, data *discordgo.InteractionCreate) {
	for _, AstridCommand := range commands.RegisteredCommands {
		if data.ApplicationCommandData().Name != AstridCommand.Command.Name {
			continue
		}

		go AstridCommand.Handler(s, data)
		return
	}
}
