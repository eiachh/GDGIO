package commands

import "github.com/bwmarrin/discordgo"

var appCommand = &discordgo.ApplicationCommand{
	Name:        "mine",
	Description: "action to mine stone or ore",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "target",
			Description: "What to mine? copper,stone",
			Required:    true,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "Copper",
					Value: "Copper",
				},
				{
					Name:  "Stone",
					Value: "Stone",
				},
			},
		},
	},
}

func GetMineAppCommand() *discordgo.ApplicationCommand {
	return appCommand
}

func HandleMineAppCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {

}
