package utils

import "github.com/bwmarrin/discordgo"

const (
	messageSource = discordgo.InteractionResponseChannelMessageWithSource
)

type Interaction struct {
	Client *discordgo.Session
	Data   *discordgo.InteractionCreate
}

type response struct {
	Embeds  []*discordgo.MessageEmbed
	Message string
	Flags   int
}

func (a *Interaction) sendResponse(data *response) error {
	error := a.Client.InteractionRespond(a.Data.Interaction, &discordgo.InteractionResponse{
		Type: messageSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:  data.Embeds,
			Content: data.Message,
			Flags:   discordgo.MessageFlags(data.Flags),
		},
	})

	return error
}

// --> Embeds

func (a *Interaction) SendEmbed(embed *discordgo.MessageEmbed, ephemeral bool) error {
	data := &response{
		Embeds: []*discordgo.MessageEmbed{embed},
	}
	if ephemeral {
		data.Flags = 1 << 6
	}
	return a.sendResponse(data)
}

// --> Messages

func (a *Interaction) SendMessage(message string, ephemeral bool) error {
	data := &response{
		Message: message,
	}
	if ephemeral {
		data.Flags = 1 << 6
	}
	return a.sendResponse(data)
}
