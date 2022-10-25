package readimage

import (
	"image"
	"image/png"
	"os"

	"github.com/pkg/errors"
)

func readPngImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer file.Close()

	pngImage, err := png.Decode(file)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return pngImage, nil
}
