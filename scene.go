package main

import (
	"context"
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	frame int64
	bg    *sdl.Texture
	birds []*sdl.Texture
}

func newScene(r *sdl.Renderer) (*Scene, error) {
	bg, err := img.LoadTexture(r, "assets/imgs/background.png")
	if err != nil {
		return nil, fmt.Errorf("could not load background image: %v", err)
	}

	var birds []*sdl.Texture
	for i := 1; i <= 4; i++ {
		bird, err := img.LoadTexture(r, fmt.Sprintf("assets/imgs/bird_frame_%d.png", i))
		if err != nil {
			return nil, fmt.Errorf("could not load birad_frame_1 image: %v", err)
		}
		birds = append(birds, bird)
	}

	return &Scene{
		bg:    bg,
		birds: birds,
	}, nil
}

func (s *Scene) run(ctx context.Context, r *sdl.Renderer) <-chan error {
	errc := make(chan error)
	ticker := time.NewTicker(10 * time.Millisecond)

	go func() {
		defer ticker.Stop()
		defer close(errc)
		for range ticker.C {
			select {
			case <-ctx.Done():
				break
			default:
				if err := s.paint(r); err != nil {
					errc <- err
				}
			}
		}
	}()

	return errc
}

func (s *Scene) paint(r *sdl.Renderer) error {
	s.frame++
	r.Clear()

	if err := r.Copy(s.bg, nil, nil); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	i := (s.frame / 10) % int64(len(s.birds))
	rect := &sdl.Rect{X: 100, Y: 100, W: 50, H: 43}
	if err := r.Copy(s.birds[i], nil, rect); err != nil {
		return fmt.Errorf("could not copy bird: %v", err)
	}

	r.Present()
	return nil
}

func (s *Scene) destroy() {
	s.bg.Destroy()
}
