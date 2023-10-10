package main

import (
	"GDCIO/commands"
	forage "GDCIO/commands/forage"
	info "GDCIO/commands/info"
	mine "GDCIO/commands/mine"
	move "GDCIO/commands/move"
	reg "GDCIO/commands/register"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var BotChannel string = "947971806617272390"
var BotGuild string = "652300397859569694"

var dg *discordgo.Session

func main() {
	// Create a new Discord session
	var token string
	tmp, _ := ioutil.ReadFile("token.txt")
	token = string(tmp)

	tmpDG, err := discordgo.New("Bot " + token)
	dg = tmpDG
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// Register a callback for the MessageCreate event
	dg.AddHandler(onReady)
	dg.AddHandler(onInteractionCreate)

	// Open a connection to Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection to Discord:", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	// Cleanly close down the Discord session
	dg.Close()
}

func registerCommands(s *discordgo.Session) {
	for _, command := range commands.GetCommands() {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, BotGuild, command)
		if err != nil {
			fmt.Println("Error creating command:", err)
		}
	}
}

func onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		command := i.ApplicationCommandData()
		switch command.Name {
		case "move":
			move.HandleMoveAppCommand(s, i)
		case "mine":
			mine.HandleMineAppCommand(s, i)
		case "forage":
			forage.HandleForageAppCommand(s, i)
		case "register":
			reg.HandleRegisterAppCommand(s, i)
		case "listactions":
			info.HandleListActionsAppCommand(s, i)
		case "char":
			info.HandleListCharStatsAppCommand(s, i)
		}

	}
}

func onReady(s *discordgo.Session, event *discordgo.Ready) {
	registerCommands(dg)
}

func SayHelloOnSelectedChannel(s *discordgo.Session, event *discordgo.Ready) {
	var targetChannel *discordgo.Channel
	channels, _ := s.GuildChannels(BotGuild)
	for _, channel := range channels {
		if channel.ID == BotChannel {
			targetChannel = channel
			break
		}
	}

	if targetChannel == nil {
		fmt.Println("Target channel not found")
		return
	}

	// Send a message to the target channel
	_, err := s.ChannelMessageSend(targetChannel.ID, "Hello, World!")
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
}

func ListOutGuildAndChannel(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Bot connected to the following servers:")

	for _, guild := range s.State.Guilds {
		fmt.Printf("- %s (ID: %s)\n", guild.Name, guild.ID)

		channels, err := s.GuildChannels(guild.ID)
		if err != nil {
			fmt.Println("Error getting channels:", err)
			continue
		}

		fmt.Println("Channels in the server:")
		for _, channel := range channels {
			fmt.Printf("- %s (ID: %s)\n", channel.Name, channel.ID)
		}
	}
}
