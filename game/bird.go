package game

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const gravity = 0.15
const birdSpriteNumber = 4

var birdHeight int32
var maxBirdY float64
var initialBirdY float64

type bird struct {
	texture *sdl.Texture
	time    uint
	y       float64
	speed   float64
	width   int32
	height  int32
}

func newBird(r *sdl.Renderer) (*bird, error) {
	t, err := img.LoadTexture(r, "assets/bird.png")
	if err != nil {
		return nil, errors.Wrap(err, "could not load bird.png")
	}

	width, height, err := getTextureSize(t)
	if err != nil {
		return nil, errors.Wrap(err, "could not get bird.png size")
	}

	maxBirdY = float64(windowHeight - birdHeight/2)
	initialBirdY = float64(windowHeight / 2)

	return &bird{
		texture: t,
		time:    0,
		y:       initialBirdY,
		speed:   0,
		width:   width,
		height:  height / birdSpriteNumber,
	}, nil
}

func (b *bird) jump() {
	b.speed = -5
}

func (b *bird) restart() {
	b.y = initialBirdY
	b.speed = 0
}

// return value means: is game over
func (b *bird) update() bool {
	b.time++
	b.y += b.speed
	b.speed += gravity

	return b.y >= maxBirdY
}

func (b *bird) paint(r *sdl.Renderer) error {
	spriteIndex := int32(b.time / 10 % birdSpriteNumber)
	y := spriteIndex * b.height
	srcRect := &sdl.Rect{X: 0, W: b.width, Y: y, H: b.height}
	dstRect := &sdl.Rect{X: 100, Y: int32(b.y) - b.height/2, W: b.width, H: b.height}
	if err := r.Copy(b.texture, srcRect, dstRect); err != nil {
		return errors.Wrap(err, "could not copy bird")
	}
	return nil
}

func (b *bird) destroy() {
	b.texture.Destroy()
}
