package game

import "github.com/pkg/errors"
import "github.com/veandco/go-sdl2/sdl"
import "github.com/veandco/go-sdl2/img"

type land struct {
	texture *sdl.Texture
	cursor  float64
	width   int32
	height  int32
}

func newLand(r *sdl.Renderer) (*land, error) {
	texture, err := img.LoadTexture(r, "assets/land.png")
	if err != nil {
		return nil, errors.Wrap(err, "could not load land.png")
	}

	width, height, err := getTextureSize(texture)
	if err != nil {
		return nil, errors.Wrap(err, "could not get land.png size")
	}

	return &land{texture, 0, width, height}, nil
}

func (l *land) update() {
	l.cursor += 0.002
	if l.cursor > 1 {
		l.cursor = 0
	}
}

func (l *land) paint(r *sdl.Renderer) error {
	// first section
	startX := (int32)(l.cursor * (float64)(l.width))
	src := &sdl.Rect{X: startX, W: l.width - startX, Y: 0, H: landHeight}
	targetWidth := int32((float64)(windowWidth) * (1 - l.cursor))
	dst := &sdl.Rect{X: 0, W: targetWidth, Y: skyHeight, H: landHeight}
	if err := r.Copy(l.texture, src, dst); err != nil {
		return errors.Wrap(err, "could not copy land")
	}

	// second section
	src = &sdl.Rect{X: 0, W: startX, Y: 0, H: landHeight}
	dst = &sdl.Rect{X: targetWidth, W: windowWidth - targetWidth, Y: skyHeight, H: landHeight}
	if err := r.Copy(l.texture, src, dst); err != nil {
		return errors.Wrap(err, "could not copy land")
	}

	return nil
}

func (l *land) destroy() {
	l.texture.Destroy()
}
