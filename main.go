package main

import(
	"github.com/era/covid-telegram/config"
	"github.com/era/covid-telegram/controller"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"strings"
	"time"
)

func main() {
	telegramBot, err := tb.NewBot(tb.Settings{
		Token:  config.GetToken(),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	router(telegramBot, routes())

	telegramBot.Start()
}


func routes() map[string]func(m *tb.Message, b *tb.Bot) {
	return map[string]func(m *tb.Message, b *tb.Bot){
		"/start": func(m *tb.Message, b *tb.Bot) {
			b.Send(m.Sender, controller.Instructions())
		},
		"/all": func(m *tb.Message, b *tb.Bot) {
			b.Send(m.Sender, controller.All())
		},
		"/country": func(m *tb.Message, b *tb.Bot) {
			country := strings.Replace(m.Text, "/country ","", 1)
			b.Send(m.Sender, controller.Country(country))
		},
	}
}

func router (telegramBot *tb.Bot, routes map[string]func(message *tb.Message, telegramBot *tb.Bot)) {
	for k, v := range routes {
		// Creates local copy
		function := v
		command := k
		telegramBot.Handle(command, func(m *tb.Message) {
			log.Print(m.Text)
			function(m, telegramBot)
		})
	}
}