package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func massagecreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	prifix := '!'

	if m.Author.ID == s.State.User.ID {
		return
	} // this gonna check the id of the masseg author is it match with the id of the bot id it self ; if yes the msg not allaw

	if len(m.Content) == 0 || m.Content[0] != byte(prifix) {
		return
	} // if msg is empty
	// if not prifix match then they gonna return the msg

	if m.Content == "!ping" {

		latancy := float64(s.HeartbeatLatency().Microseconds()) / 1000.0
		pintmsg := fmt.Sprintf("🏓 Pong! : **%.1fms**", latancy)
		s.ChannelMessageSend(m.ChannelID, pintmsg)

	}

	// new command

	if strings.HasPrefix(m.Content, "!ask ") {

		// removeing the command and prefix form the command

		question := strings.TrimPrefix(m.Content, "!ask ")

		// showing thinking in the channel

		s.ChannelTyping(m.ChannelID)

		// get the answer form the function we made

		answer := askGemini(question)

		// sent it by the bot

		s.ChannelMessageSend(m.ChannelID, answer)

	}

}

func askGemini(userInput string) string {

	api := os.Getenv("API")
	ctx := context.Background()

	// api key

	client, err := genai.NewClient(ctx, option.WithAPIKey(api))
	// client, _ := genai.NewClient(ctx, option.WithAPIKey(api))
	if err != nil {
		return " Failed to cannect to the brain!"

	}
	defer client.Close()

	// model := client.GenerativeModel("gemini-2.5-flash-lite")
	// model := client.GenerativeModel("gemini-3-flash-preview")
	model := client.GenerativeModel("gemini-3.1-flash-lite-preview")

	resp, err := model.GenerateContent(ctx, genai.Text(userInput))

	if err != nil {
		// This will print the actual error from Google in your Discord
		return fmt.Sprintf("❌ Captin there is a API Error: %v", err)
	}

	return fmt.Sprint(resp.Candidates[0].Content.Parts[0])

}

func main() {

	// loade the token

	godotenv.Load()

	token := os.Getenv("BOT_TOKEN")

	// create a session

	dg, _ := discordgo.New("Bot " + token)

	// set intent

	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentMessageContent)
	dg.AddHandler(massagecreate)

	// open connection

	err := dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	// keep running until ctrl+c
	fmt.Println("Bot is now online. press CTRL-C to exit. ")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()

	// fmt.println("hello world")

}
