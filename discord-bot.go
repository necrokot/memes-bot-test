package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"net/http"
	"os"
	"strings"
)

type PhotoStruct struct {
	reader io.Reader
	name   string
}

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + (os.Getenv("DISCORD_APITOKEN")))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if len(m.Content) > 0 {
		sendText(m.Content)
	}

	if len(m.Attachments) == 1 {
		res, _ := http.DefaultClient.Get(m.Attachments[0].URL)
		defer res.Body.Close()
		if strings.HasPrefix(m.Attachments[0].ContentType, "image") {
			sendPhoto(res.Body, m.Attachments[0].Filename)
		} else {
			sendDocument(res.Body, m.Attachments[0].Filename)
		}
	} else if len(m.Attachments) > 1 {

		photoGroup := make([]PhotoStruct, len(m.Attachments))

		for i, attachment := range m.Attachments {
			res, _ := http.DefaultClient.Get(attachment.URL)
			photoGroup[i].reader = res.Body
			defer res.Body.Close()
			photoGroup[i].name = attachment.Filename
		}

		sendPhotoGroup(photoGroup)
	}
}

func runDiscordBot() {
	Start()

	<-make(chan struct{})
	return
}
