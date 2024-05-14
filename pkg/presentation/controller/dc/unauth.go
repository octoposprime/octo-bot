package presentation

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (a *DcAPI) PingHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {

	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "pong!",
			},
		})
	}

}

func (a *DcAPI) CalculateHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {

	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		options := i.ApplicationCommandData().Options
		value := options[0].IntValue()

		resultValue, err := a.queryHandler.Calculate(context.Background(), value)
		resultMessage := ""
		if err != nil {
			resultMessage = fmt.Sprintf("Error calculating %d!", value)
		} else {
			resultMessage = fmt.Sprintf("%d! = %d", value, resultValue)
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: resultMessage,
			},
		})
	}

}
