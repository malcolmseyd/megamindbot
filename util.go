package main

import (
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
)

// the meme pattern should match on anything after "no" and before punctuation
// trim any trailing whitespace, and don't let the body just be whitespace
// "Regex is a write-only language" - Bill Bird
var memePattern = regexp.MustCompile(`[nN][oO]\s+([^.,!?;:\s][^.,!?;:]*)\s*`)

// format text so it looks like the meme
func memeify(text string) string {
	return "NO " + strings.ToUpper(strings.TrimSpace(text)) + "?"
}

// block until Ctrl+C
func blockUntilInterrupt() {
	c := make(chan os.Signal, 1)
	go signal.Notify(c, os.Interrupt, syscall.SIGINT)
	<-c
}

func mustEnv(key string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		log.Fatalln("missing required environment variable:", key)
	}
	return
}
