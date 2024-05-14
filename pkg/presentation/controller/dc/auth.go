package presentation

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (a *DcAPI) UsersHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {

	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		options := i.ApplicationCommandData().Options
		filter := options[0].StringValue()

		resultValue, err := a.queryHandler.Users(context.Background(), filter)
		resultMessage := ""
		if err != nil {
			resultMessage = fmt.Sprintf("Error getting users %s!", filter)
		} else {
			resultMessage = fmt.Sprintf("Users: %v", resultValue)
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: resultMessage,
			},
		})
	}

}
