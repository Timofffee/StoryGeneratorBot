package main

import (
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("YOUR_TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	insertPart := []string{"Baka", "Jhonny", "Voody", "Gregory"}
	templates := []string{"##insert sit on chair", "I love ##insert and ##insert!!", "##insert killed ##insert"}

	insertPartReg := regexp.MustCompile(`##insert\b`)

	for update := range updates {
		if update.Message != nil {
			if strings.HasPrefix(update.Message.Text, "/gen") {
				rand.Seed(time.Now().UnixNano())

				text := replace(templates[rand.Intn(len(templates))], insertPartReg, insertPart)

				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	}
}

func replace(text string, reg *regexp.Regexp, textArr []string) string {
	s := reg.ReplaceAllStringFunc(text, func(x string) string {
		var r string = textArr[rand.Intn(len(textArr))]
		return r
	})

	return s
}
