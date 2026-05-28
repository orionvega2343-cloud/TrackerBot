package main

import (
	"TrackerBot/Internal/config"
	"TrackerBot/Internal/db"
	"TrackerBot/Internal/handlers"
	"TrackerBot/Internal/repository"
	"TrackerBot/Internal/service"
	"log"
	"net/http"
	"net/url"
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
	httpClient := &http.Client{}
	if cfg.ProxyURL != "" {
		proxyURL, err := url.Parse(cfg.ProxyURL)
		if err != nil {
			log.Fatal("invalid PROXY_URL:", err)
		}
		httpClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}

	pref := tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		Client: httpClient,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	hRepo := repository.NewHabitsRepo(database)
	cRepo := repository.NewCompleteRepo(database)

	hService := service.NewHabitsService(hRepo)
	cService := service.NewCompleteService(cRepo)

	hHandler := handlers.NewHabitsHandler(hService, b)
	cHandler := handlers.NewCompleteHabitsHandler(cService, b)

	hHandler.Reg()
	cHandler.Reg()

	b.Start()
}
