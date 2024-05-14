package presentation

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	pp_command "github.com/octoposprime/octo-bot/internal/application/presentation/port/command"
	pp_query "github.com/octoposprime/octo-bot/internal/application/presentation/port/query"
	dto "github.com/octoposprime/octo-bot/pkg/presentation/dto"
	tconfig "github.com/octoposprime/octo-bot/tool/config"
)

// DcAPI is the Dc API for the application
type DcAPI struct {
	queryHandler   pp_query.QueryPort
	commandHandler pp_command.CommandPort

	DcCommandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
	Commands          []*dto.Command
}

// NewDc creates a new instance of Dc
func NewDcAPI(qh pp_query.QueryPort, ch pp_command.CommandPort) *DcAPI {
	api := &DcAPI{
		queryHandler:      qh,
		commandHandler:    ch,
		DcCommandHandlers: make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)),
		Commands:          make([]*dto.Command, 0),
	}
	return api
}

// Serve starts the API server
func (a *DcAPI) Start(dcConfig tconfig.DcConfig) {
	dg, err := discordgo.New("Bot " + dcConfig.Dc.Token)
	if err != nil {
		panic(err)
	}

	err = dg.Open()
	if err != nil {
		panic(err)
	}
	defer dg.Close()

	a.DcCommandHandlers["ping"] = a.PingHandler()
	a.DcCommandHandlers["calculate"] = a.CalculateHandler()
	a.DcCommandHandlers["users"] = a.UsersHandler()

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if cmd, ok := a.DcCommandHandlers[i.ApplicationCommandData().Name]; ok {
			fmt.Printf("Received command: %s\n", i.ApplicationCommandData().Name)
			cmd(s, i)
		}
	})

	a.Commands = append(a.Commands, &dto.Command{
		Name:        "ping",
		Description: "Replies with pong!",
		Options:     nil,
	})

	a.Commands = append(a.Commands, &dto.Command{
		Name:        "calculate",
		Description: "Factorial calculation",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "value",
				Description: "Calculate the factorial of this number",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    true,
			},
		},
	})

	a.Commands = append(a.Commands, &dto.Command{
		Name:        "users",
		Description: "Get users by filter",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "filter",
				Description: "Filter to search users",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	})

	for _, v := range a.Commands {
		_, err := dg.ApplicationCommandCreate(dg.State.User.ID, dcConfig.Dc.GuildID, v.ToApplicationCommand())
		if err != nil {
			fmt.Printf("Cannot create '%v' command: %v\n", v.Name, err)
		}
	}

	fmt.Println("Bot is now running. Send 'close' before exit.")
	var msg string
	fmt.Scanf("%s", &msg)
	if msg == "close" {
		for _, v := range a.Commands {
			err := dg.ApplicationCommandDelete(dg.State.User.ID, dcConfig.Dc.GuildID, v.ToApplicationCommand().ID)
			if err != nil {
				fmt.Printf("Cannot delete '%v' command: %v\n", v.Name, err)
			}
		}
	}
	fmt.Println("Bot is now closing.")
}
