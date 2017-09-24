package game

import "github.com/veandco/go-sdl2/sdl"

// func paintTiled(r *sdl.Renderer, txt *sdl.Texture, dst *sdl.Rect) error {
// 	width, height, err := getTextureSize(txt)
// 	if err != nil {
// 		return errors.Wrap(err, "could not get texture size")
// 	}

// 	return nil
// }

func getTextureSize(txt *sdl.Texture) (int32, int32, error) {
	_, _, width, height, err := txt.Query()
	if err != nil {
		return 0, 0, err
	}

	return width, height, nil
}
