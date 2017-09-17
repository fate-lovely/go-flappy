package game

import (
	"fmt"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Game struct {
	events   chan event
	window   *sdl.Window
	renderer *sdl.Renderer
	scene    *scene
	gameOver bool
}

type event int

const (
	eventNone    event = 0
	eventQuit    event = iota
	eventJump    event = iota
	eventRestart event = iota
)

var windowWidth int
var windowHeight int

func NewGame(width, height int) *Game {
	windowWidth = width
	windowHeight = height
	maxY = float64(windowHeight - birdHeight/2)
	initialY = float64(windowHeight / 2)

	return &Game{
		events: make(chan event),
	}
}

// must be called in main goroutine
func (g *Game) Run() error {
	g.init()
	defer g.destroy()

	runtime.LockOSThread()

	errc := g.loop()

	for {
		select {
		case err := <-errc:
			return err
		default:
			evt := sdl.WaitEvent()
			switch evt.(type) {
			case *sdl.QuitEvent:
				g.events <- eventQuit
			case *sdl.MouseButtonEvent:
				g.events <- eventJump
			}
		}
	}
}

/*----------  Private Methods  ----------*/

func (g *Game) loop() chan error {
	errc := make(chan error)

	go func() {
		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()
		defer close(errc)
		for range ticker.C {
			evt := g.fetchEvent()
			if evt == eventQuit {
				return
			}
			if err := g.tick(evt); err != nil {
				errc <- err
			}
		}
	}()

	return errc
}

func (g *Game) fetchEvent() event {
	select {
	case evt := <-g.events:
		return evt
	default:
		return eventNone
	}
}

func (g *Game) restart() {
	g.gameOver = false
	g.scene.restart()
	// clear all events
outer:
	for {
		select {
		case <-g.events:
		default:
			break outer
		}
	}
}

func (g *Game) tick(evt event) error {
	if evt == eventRestart {
		g.restart()
		return g.paint()
	}

	g.update(evt)
	return g.paint()
}

func (g *Game) paint() error {
	if g.gameOver {
		return g.drawTitle("Game Over")
	}
	return g.scene.paint(g.renderer)
}

func (g *Game) update(evt event) {
	if g.gameOver {
		return
	}

	g.gameOver = g.scene.update(evt)
	if g.gameOver {
		time.AfterFunc(2*time.Second, func() {
			g.events <- eventRestart
		})
	}
}

func (g *Game) init() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("could not init sdl: %v", err)
	}

	w, r, err := sdl.CreateWindowAndRenderer(windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("could not create window: %v", err)
	}
	g.renderer = r
	g.window = w

	// scene
	scene, err := newScene(g.renderer)
	if err != nil {
		return err
	}
	g.scene = scene

	return nil
}

func (g *Game) destroy() {
	sdl.Quit()
	g.window.Destroy()
	g.renderer.Destroy()
	g.scene.destroy()
}

func (g *Game) drawTitle(text string) error {
	r := g.renderer

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
	surface, err := font.RenderUTF8_Solid(text, color)
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
