package main

import (
	"os"

	"github.com/awakot/go-player/internal/player"
)

func main() {

	if e := run(); e != nil {
		os.Exit(1)
	}
}

func run() error {

	app := player.Player{}
	app.New()

	app.LoadMusic()
	return nil
}
