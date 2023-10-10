package commands

import (
	forage "GDCIO/commands/forage"
	info "GDCIO/commands/info"
	mine "GDCIO/commands/mine"
	move "GDCIO/commands/move"
	reg "GDCIO/commands/register"

	"github.com/bwmarrin/discordgo"
)

var commands = []*discordgo.ApplicationCommand{
	reg.GetRegisterAppCommand(),
	mine.GetMineAppCommand(),
	move.GetMoveAppCommand(),
	info.GetListActionsAppCommand(),
	info.GetListCharStatsAppCommand(),
	forage.GetForageAppCommand(),
}

func GetCommands() []*discordgo.ApplicationCommand {
	return commands
}
