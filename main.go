package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const prefix string = ""

func main() {
	sess, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("Error creating session: ", err)
		return
	}

	sess.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is running!")
	})

	err = sess.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	defer sess.Close()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	sess.Close()
}
