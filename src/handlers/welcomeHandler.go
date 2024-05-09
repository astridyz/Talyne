package handlers

import (
	"fmt"

	aura "github.com/astridyz/Aura/src"

	"github.com/bwmarrin/discordgo"
)

const (
	GuildWelcomeID = "1235669274622820362"
	WelcomeChannel = "1235669274622820365"
)

func WelcomeHandler(s *discordgo.Session, data *discordgo.GuildMemberAdd) {

	if data.GuildID != "1235669274622820362" {
		return
	}

	_, error := s.ChannelMessageSend(WelcomeChannel, fmt.Sprintf("Welcome to Astrid's Land, %v!", data.Member.User.Mention()))
	if error != nil {
		aura.Panicf("Error on sending message: %v", error)
	}
}
