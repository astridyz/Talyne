package utils

import "github.com/bwmarrin/discordgo"

func CreateErrorEmbed(message string) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{Name: "❌ Error!"},
		Description: ":construction: " + message,
		Color:       0xD95B5B,
	}

	return embed
}
