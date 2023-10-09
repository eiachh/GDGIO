package connection

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func HandleHttp(s *discordgo.Session, i *discordgo.InteractionCreate, payload []byte, url string) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}

	tmp, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Print(string(tmp))

	defer resp.Body.Close()
	//message := "definitly not implemented"
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: string(tmp),
		},
	})
}
