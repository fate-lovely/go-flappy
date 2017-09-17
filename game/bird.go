package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const gravity = 0.15
const birdWidth = 50
const birdHeight = 43

var maxY float64
var initialY float64

type bird struct {
	textures []*sdl.Texture
	time     int64
	y        float64
	speed    float64
}

func newBird(r *sdl.Renderer) (*bird, error) {
	var textures []*sdl.Texture
	for i := 1; i <= 4; i++ {
		t, err := img.LoadTexture(r, fmt.Sprintf("assets/imgs/bird_frame_%d.png", i))
		if err != nil {
			return nil, fmt.Errorf("could not load birad frame %d image: %v", i, err)
		}
		textures = append(textures, t)
	}

	return &bird{
		textures: textures,
		time:     0,
		y:        initialY,
		speed:    0,
	}, nil
}

func (b *bird) jump() {
	b.speed = -5
}

func (b *bird) restart() {
	b.y = initialY
	b.speed = 0
}

// return value means: is game over
func (b *bird) update() bool {
	b.time++
	b.y += b.speed
	b.speed += gravity

	return b.y >= maxY
}

func (b *bird) paint(r *sdl.Renderer) error {
	i := (b.time / 10) % int64(len(b.textures))
	rect := &sdl.Rect{X: 100, Y: int32(b.y - birdHeight/2), W: 50, H: 43}
	if err := r.Copy(b.textures[i], nil, rect); err != nil {
		return fmt.Errorf("could not copy bird: %v", err)
	}
	return nil
}

func (b *bird) destroy() {
	for _, t := range b.textures {
		t.Destroy()
	}
}
