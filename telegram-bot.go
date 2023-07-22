package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"os"
)

func runTelegramBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(-1001978215979, update.Message.Text)
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}

}

func sendText(txt string) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	msg := tgbotapi.NewMessage(-1001978215979, txt)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func sendDocument(f io.Reader, name string) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	file := tgbotapi.FileReader{
		Name:   name,
		Reader: f,
	}

	docConfig := tgbotapi.NewDocument(-1001978215979, file)
	if _, err := bot.Send(docConfig); err != nil {
		panic(err)
	}
}

func sendPhoto(ph io.Reader, name string) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	photo := tgbotapi.FileReader{
		Name:   name,
		Reader: ph,
	}

	phConfig := tgbotapi.NewPhoto(-1001978215979, photo)
	if _, err := bot.Send(phConfig); err != nil {
		panic(err)
	}
}

func sendPhotoGroup(photoGroup []PhotoStruct) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	photoGroupFileReader := make([]tgbotapi.FileReader, len(photoGroup))

	for i := range photoGroup {
		photoGroupFileReader[i] = tgbotapi.FileReader{
			Name:   photoGroup[i].name,
			Reader: photoGroup[i].reader,
		}
	}

	photoGroupInputMedia := make([]interface{}, len(photoGroupFileReader))
	for i := range photoGroupFileReader {
		photoGroupInputMedia[i] = tgbotapi.NewInputMediaPhoto(photoGroupFileReader[i])
	}

	photoGroupConfig := tgbotapi.NewMediaGroup(-1001978215979, photoGroupInputMedia)
	if _, err := bot.SendMediaGroup(photoGroupConfig); err != nil {
		panic(err)
	}
}
