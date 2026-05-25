package main

import (
	"TrackerBot/Internal/config"
	"TrackerBot/Internal/db"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	cfg := config.MustLoadConfig()
	database, err := db.ConnDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	//Initialize bot
	pref := tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	b.Start()
}
