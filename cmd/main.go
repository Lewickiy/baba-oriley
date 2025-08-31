package main

import (
	"baba-oriley/internal/player"
	"log"
)

func main() {
	events, err := player.LoadEvents("demo")
	if err != nil {
		log.Fatal(err)
	}

	if err := player.PlayEvents(events, 44100, "baba", 1.9); err != nil {
		log.Fatal(err)
	}
}
