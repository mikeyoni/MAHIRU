package main

import (
	// "bytes"
	"fmt"
	"os"
	"os/signal"

	// "strings"
	"syscall"

	// "os/exec"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
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

		latancy := s.HeartbeatLatency()
		pintmsg := fmt.Sprintf("🏓 Pong! : **%v**", latancy)
		s.ChannelMessageSend(m.ChannelID, pintmsg)

	}
}

func main() {

	// loade the token

	godotenv.Load()

	token := os.Getenv("BOT_TOKEN")

	// create a session

	dg, _ := discordgo.New("Bot " + token)

	// set intent

	dg.Identify.Intents = discordgo.IntentsGuildMessages

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
