package commands

import (
	"GDCIO/commons"
	"GDCIO/connection"
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

var appCommandListCharStat = &discordgo.ApplicationCommand{
	Name:        "char",
	Description: "List all the possible actions your character can do.",
}

func GetListCharStatsAppCommand() *discordgo.ApplicationCommand {
	return appCommandListCharStat
}

func HandleListCharStatsAppCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	basicCommand := commons.BasicCommand{
		OwnerId:  i.Interaction.Member.User.ID,
		Command:  "listcharstats",
		ExtraArg: "",
	}
	payload, _ := json.Marshal(basicCommand)
	url := "http://127.0.0.1:8080/command/info"
	connection.HandleHttp(s, i, payload, url)
}
