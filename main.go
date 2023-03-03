package main

import (
	"flag"
	tgClient "go_projects/tg_bot/clients/telegram"
	event_consumer "go_projects/tg_bot/consumer/event-consumer"
	"go_projects/tg_bot/events/telegram"
	"go_projects/tg_bot/storage/files"
	"log"
)

const (
	tgBotHost   = "core.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsPricessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsPricessor, eventsPricessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal()
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access")
	flag.Parse()

	if *token == "" {
		log.Fatal()
	}

	return *token
}
