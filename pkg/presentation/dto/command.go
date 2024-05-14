package presentation

import "github.com/bwmarrin/discordgo"

type Command struct {
	Name        string
	Description string
	Options     []*discordgo.ApplicationCommandOption
}

func (c *Command) ToApplicationCommand() *discordgo.ApplicationCommand {

	return &discordgo.ApplicationCommand{
		Name:        c.Name,
		Description: c.Description,
		Options:     c.Options,
	}

}
