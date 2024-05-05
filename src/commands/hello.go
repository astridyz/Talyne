package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/astridyz/talyne-discord-bot/utils"
)

var Hello_Command = &Command{
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
	var Interaction = utils.Interaction{Client: s, Data: data}

	if Options == nil || Options[0] == nil {
		Interaction.SendEmbed(utils.CreateErrorEmbed("No user have been sent."), true)
		return
	}

	Interaction.SendMessage(
		fmt.Sprintf("Hello, nice to meet you, %v!", Options[0].UserValue(s).Mention()),
		false,
	)

}
