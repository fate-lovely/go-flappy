package main

import (
	"fmt"
	"os"

	"github.com/fate-lovely/go-flappy/game"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

func main() {
	game := game.NewGame(windowWidth, windowHeight)

	if err := game.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
