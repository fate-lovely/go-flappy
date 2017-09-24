package game

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	sky  *sdl.Texture
	land *land
	bird *bird
}

func newScene(r *sdl.Renderer) (*scene, error) {
	sky, err := img.LoadTexture(r, "assets/sky.png")
	if err != nil {
		return nil, errors.Wrap(err, "could not load sky.png")
	}

	land, err := newLand(r)
	if err != nil {
		return nil, err
	}

	bird, err := newBird(r)
	if err != nil {
		return nil, err
	}

	return &scene{
		sky:  sky,
		land: land,
		bird: bird,
	}, nil
}

// return value means: is game over
func (s *scene) update(evt event) bool {
	if evt == eventJump {
		s.bird.jump()
	}

	s.land.update()

	return s.bird.update()
}

func (s *scene) restart() {
	s.bird.restart()
}

func (s *scene) paint(r *sdl.Renderer) error {
	r.Clear()

	skyDst := &sdl.Rect{X: 0, Y: 0, W: windowWidth, H: skyHeight}
	if err := r.Copy(s.sky, nil, skyDst); err != nil {
		return errors.Wrap(err, "could not copy sky")
	}

	if err := s.land.paint(r); err != nil {
		return err
	}

	if err := s.bird.paint(r); err != nil {
		return err
	}

	r.Present()

	return nil
}

func (s *scene) destroy() {
	s.sky.Destroy()
	s.land.destroy()
	s.bird.destroy()
}
