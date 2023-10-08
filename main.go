package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var BotChannel string = "947971806617272390"
var BotGuild string = "652300397859569694"

var dg *discordgo.Session

var commands = []*discordgo.ApplicationCommand{
	{
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
	},
	{
		Name:        "mine",
		Description: "action to mine stone or ore",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "target",
				Description: "What to mine? copper,stone",
				Required:    true,
			},
		},
	},
	{
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
	},
}

type RegisterCommand struct {
	OwnerId    string `json:"OwnerId"`
	PlayerName string `json:"PlayerName"`
}

func main() {
	// Create a new Discord session
	token := "OTQ3OTY5MjQwNjAwODM0MTE4.GN74y-.H1hs_p7Cs9TTRacd84gH-LGTMSctG0bZqEK7fw"

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
	for _, command := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, BotGuild, command)
		if err != nil {
			fmt.Println("Error creating command:", err)
		}
	}
}

func onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		command := i.ApplicationCommandData()

		if command.Name == "move" {
			message := "Trying to move: " + command.Options[0].StringValue()
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: message,
				},
			})
		}
		if command.Name == "mine" {
			message := "definitly not implemented"
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: message,
				},
			})
		}
		if command.Name == "register" {
			regCommand := RegisterCommand{
				OwnerId:    i.Interaction.Member.User.ID,
				PlayerName: command.Options[0].StringValue(),
			}
			payload, _ := json.Marshal(regCommand)
			resp, err := http.Post("http://127.0.0.1:8080/command/register", "application/json", bytes.NewBuffer(payload))
			if err != nil {
				fmt.Println("Error making POST request:", err)
				return
			}
			defer resp.Body.Close()
			message := "definitly not implemented"
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: message,
				},
			})
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
