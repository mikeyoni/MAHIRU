package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	// "os/exec"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

)

func main(){

	// loade the token 

	godotenv.Load()

	token:= os.Getenv("BOT_TOKEN")
	

	// create a session 

	dg, _ := discordgo.New("Bot " + token)

	// set intent 

	dg.Identify.Intents = discordgo.IntentsGuildMessages 

	// open connection 

	err := dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	// keep running until ctrl+c 
	fmt.Println("Bot is now online. press CTRL-C to exit. ")
	sc :=make(chan os.Signal, 1)
	signal.Notify(sc , syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()

	// fmt.println("hello world")

}