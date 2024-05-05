package utils

import "github.com/bwmarrin/discordgo"

const (
	MessageSource = discordgo.InteractionResponseChannelMessageWithSource
)

type AstridInteraction struct {
	Client *discordgo.Session
	Data   *discordgo.InteractionCreate
}

type AstridResponse struct {
	Embeds  []*discordgo.MessageEmbed
	Message string
	Flags   int
}

func (a *AstridInteraction) sendResponse(data *AstridResponse) error {
	error := a.Client.InteractionRespond(a.Data.Interaction, &discordgo.InteractionResponse{
		Type: MessageSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:  data.Embeds,
			Content: data.Message,
			Flags:   discordgo.MessageFlags(data.Flags),
		},
	})
	return error
}

// --> Embeds
func (a *AstridInteraction) SendEphemeralEmbed(embed *discordgo.MessageEmbed) error {
	data := &AstridResponse{
		Embeds: []*discordgo.MessageEmbed{embed},
		Flags:  1 << 6,
	}
	return a.sendResponse(data)
}

func (a *AstridInteraction) SendMessageEmbed(embed *discordgo.MessageEmbed) error {
	data := &AstridResponse{
		Embeds: []*discordgo.MessageEmbed{embed},
	}
	return a.sendResponse(data)
}

// --> Messages
func (a *AstridInteraction) SendEphemeralMessage(message string) error {
	data := &AstridResponse{
		Message: message,
		Flags:   1 << 6,
	}
	return a.sendResponse(data)
}

func (a *AstridInteraction) SendMessage(message string) error {
	data := &AstridResponse{
		Message: message,
	}

	return a.sendResponse(data)
}
