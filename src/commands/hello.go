package commands

import (
	"github.com/bwmarrin/discordgo"

	"github.com/astridyz/talyne-discord-bot/utils"
)

var Hello_Command = &AstridCommand{
	Command: &discordgo.ApplicationCommand{
		Name:        "hello",
		Description: "Description",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "person",
				Description: "Who you want yo say hello!",
				// --> Required:    true,
			},
		},
	},
	Handler: helloMessageReceiver,
}

func helloMessageReceiver(s *discordgo.Session, data *discordgo.InteractionCreate) {

	Options := data.ApplicationCommandData().Options

	if Options == nil || Options[0] == nil {
		s.InteractionRespond(data.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{Embeds: []*discordgo.MessageEmbed{utils.CreateErrorEmbed("No users have been sent.")}},
		})

		return
	}

	s.InteractionRespond(data.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: "Hello, nice to meet you, " + Options[0].UserValue(s).Mention() + "!"},
	})
}
