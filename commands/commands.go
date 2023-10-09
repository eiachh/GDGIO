package commands

import (
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
}

func GetCommands() []*discordgo.ApplicationCommand {
	return commands
}
