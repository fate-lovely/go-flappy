package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("could not init sdl: %v", err)
	}
	defer sdl.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("could not create window: %v", err)
	}
	defer w.Destroy()

	if err := drawTitle(r); err != nil {
		return fmt.Errorf("could not draw title: %v", err)
	}

	time.Sleep(5 * time.Second)
	return nil
}

func drawTitle(r *sdl.Renderer) error {
	r.Clear()

	err := ttf.Init()
	if err != nil {
		fmt.Errorf("could not init ttf: %v", err)
	}

	font, err := ttf.OpenFont("assets/fonts/Flappy.ttf", 20)
	if err != nil {
		return fmt.Errorf("could not open font: %v", err)
	}
	defer font.Close()

	color := sdl.Color{R: 255, G: 100, B: 0, A: 255}
	surface, err := font.RenderUTF8_Solid("Flappy Gopher", color)
	if err != nil {
		return fmt.Errorf("could not render font: %v", err)
	}
	defer surface.Free()

	texture, err := r.CreateTextureFromSurface(surface)
	if err != nil {
		return fmt.Errorf("could not create texture: %v", err)
	}
	defer texture.Destroy()

	if err := r.Copy(texture, nil, nil); err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}

	r.Present()
	return nil
}
