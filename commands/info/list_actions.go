package commands

import (
	"GDCIO/commons"
	"GDCIO/connection"
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

var appCommand = &discordgo.ApplicationCommand{
	Name:        "listactions",
	Description: "List all the possible actions your character can do.",
}

func GetListActionsAppCommand() *discordgo.ApplicationCommand {
	return appCommand
}

func HandleListActionsAppCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	basicCommand := commons.BasicCommand{
		OwnerId:  i.Interaction.Member.User.ID,
		Command:  "listactions",
		ExtraArg: "",
	}
	payload, _ := json.Marshal(basicCommand)
	url := "http://127.0.0.1:8080/command/info"
	connection.HandleHttp(s, i, payload, url)
}
