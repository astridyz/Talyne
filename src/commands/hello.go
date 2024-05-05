package commands

import (
	"github.com/bwmarrin/discordgo"
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
				Required:    true,
			},
		},
	},
	Handler: helloMessageReceiver,
}

func helloMessageReceiver(s *discordgo.Session, data *discordgo.InteractionCreate) {

	options := data.ApplicationCommandData().Options
	/*
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))

			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			var user *discordgo.User

			if option, ok := optionMap["person"]; ok {
				user = option.UserValue(s)
			}
	*/

	s.InteractionRespond(data.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: "Hello, nice to meet you, " + options[0].UserValue(s).Mention() + "!"},
	})
}
