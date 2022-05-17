package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//handle all the messages coming and if it's a valid command run the command handler
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, conf.Prefix) {
		if m.Author.ID == conf.BotID {
			return
		}

		elements := strings.Split(m.Content, " ")
		switch elements[0] {
		case conf.Prefix + "ping":
			s.ChannelMessageSend(m.ChannelID, "pong")
		case conf.Prefix + "query", conf.Prefix + "q":
			// QueryHandler(s, m, elements[1:])
		// case conf.Prefix + "newRefresh", conf.Prefix + "nr":
		// 	UpdateHandler(s, m, elements[1:])
		// case conf.Prefix + "user", conf.Prefix + "u":
		// 	QuitHandler(s, m)
		// case conf.Prefix + "block", conf.Prefix + "b":
		// 	PBHandler(s, m)
		// case conf.Prefix + "unlock", conf.Prefix + "ul":
		// 	LeaderboardHandler(s, m)
		// case conf.Prefix + "grant", conf.Prefix + "g":
		// 	GrantHandler(s, m, elements[1:])
		// case conf.Prefix + "revoke", conf.Prefix + "r":
		// 	RevokeHandler(s, m, elements[1:])
		// case conf.Prefix + "help", conf.Prefix + "h":
		// 	HelpHandler(s, m)
		default:
			_, _ = s.ChannelMessageSend(m.ChannelID, "codice sconosciuto, usa "+conf.Prefix+"help per sapere i codici che puoi usare")
		}
	}
}

//main
func main() {
	discord, err := discordgo.New("Bot " + conf.Token)
	if err != nil {
		log.Fatal(err)
	}

	u, err := discord.User("@me")
	if err != nil {
		fmt.Println(err)
	}
	conf.BotID = u.ID

	discord.AddHandler(MessageHandler)
	err = discord.Open()
	if err != nil {
		log.Fatalf("error opening discord: %v", err.Error())
	}

	fmt.Println("wow i am working :D ")
	<-make(chan struct{})
}
