package commands

import (
	"GDCIO/commons"
	"GDCIO/connection"
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

var appCommandListCharStat = &discordgo.ApplicationCommand{
	Name:        "forage",
	Description: "Forage for food",
}

func GetForageAppCommand() *discordgo.ApplicationCommand {
	return appCommandListCharStat
}

func HandleForageAppCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	basicCommand := commons.BasicCommand{
		OwnerId:  i.Interaction.Member.User.ID,
		Command:  "forage",
		ExtraArg: "",
	}
	payload, _ := json.Marshal(basicCommand)
	url := "http://127.0.0.1:8080/command/basic"
	connection.HandleHttp(s, i, payload, url)
}
