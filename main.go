package main

import (
	"context"
	"log"

	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
)

// required environment variables
var username = mustEnv("IMGFLIP_USERNAME")
var password = mustEnv("IMGFLIP_PASSWORD")
var token = mustEnv("BOT_TOKEN")

// the bot should get Guild messages as well as DMs
var botIntents = []gateway.Intents{gateway.IntentGuildMessages, gateway.IntentDirectMessages}

// this controls the bot
// being global is okay since nobody uses it until it's assigned
var s *session.Session

func main() {
	// setup bot
	s = session.NewWithIntents("Bot "+token, botIntents...)
	s.AddHandler(handleMessage)

	// connect bot to Discord
	if err := s.Open(context.Background()); err != nil {
		log.Fatalln("Failed to connect:", err)
	}
	defer s.Close()

	// block until Ctrl+C
	blockUntilInterrupt()
}
