package commands

import (
	"GDCIO/commons"
	"GDCIO/connection"
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

var appCommand = &discordgo.ApplicationCommand{
	Name:        "register",
	Description: "Register if you did not play yet",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "playername",
			Description: "The name of your character",
			Required:    true,
		},
	},
}

func GetRegisterAppCommand() *discordgo.ApplicationCommand {
	return appCommand
}

func HandleRegisterAppCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	command := i.ApplicationCommandData()

	regCommand := commons.RegisterCommand{
		OwnerId:    i.Interaction.Member.User.ID,
		PlayerName: command.Options[0].StringValue(),
	}
	payload, _ := json.Marshal(regCommand)
	url := "http://127.0.0.1:8080/command/register"
	connection.HandleHttp(s, i, payload, url)
}
