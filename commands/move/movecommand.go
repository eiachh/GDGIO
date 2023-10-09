package commands

import (
	"GDCIO/commons"
	"GDCIO/connection"
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

var appCommand = &discordgo.ApplicationCommand{
	Name:        "move",
	Description: "Make the bot say something",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "direction",
			Description: "North, East, South, West",
			Required:    true,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "North",
					Value: "North",
				},
				{
					Name:  "East",
					Value: "East",
				},
				{
					Name:  "South",
					Value: "South",
				},
				{
					Name:  "West",
					Value: "West",
				},
			},
		},
	},
}

func GetMoveAppCommand() *discordgo.ApplicationCommand {
	return appCommand
}

func HandleMoveAppCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	command := i.ApplicationCommandData()

	moveCommand := commons.BasicCommand{
		OwnerId:  i.Interaction.Member.User.ID,
		Command:  "Move",
		ExtraArg: command.Options[0].StringValue(),
	}
	payload, _ := json.Marshal(moveCommand)
	url := "http://127.0.0.1:8080/command/basic"
	connection.HandleHttp(s, i, payload, url)
}
