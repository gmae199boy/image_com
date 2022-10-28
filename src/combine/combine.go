package combine

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type CombineOptions struct {
	W int
	H int
	Images []*image.Image
	OutputPath string
	ImageName string
	SaveAs ImageType
}

type ImageType int

const (
	Png ImageType = iota
	Jpeg
)

func Image(opt *CombineOptions) error {
	bounds := image.Rect(0, 0, opt.W, opt.H)
	bg := image.NewRGBA(bounds)
	
	for _, img := range opt.Images {
		draw.Draw(bg, bounds, *img, image.Point{}, draw.Over)
	}

	switch opt.SaveAs {
	case Png: 
		if err := saveImageAsPng(bg, opt.OutputPath, opt.ImageName); err != nil {
			return err
		}
	case Jpeg: 
		if err := saveImageAsJpeg(bg, opt.OutputPath, opt.ImageName); err != nil {
			return err
		}
	default: 
		return errors.New("not set save as type")
	}

	return nil
}

func saveImageAsPng(img *image.RGBA, out string, imageName string) error {
	result, err := os.Create(filepath.Join(out, imageName + ".png"))
	if err != nil {
		return errors.New("fail to create file")
	}

	if err := png.Encode(result, img); err != nil {
		result.Close()
		return errors.New(err.Error())
	}

	if err := result.Close(); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func saveImageAsJpeg(img *image.RGBA, out string, imageName string) error {
	result, err := os.Create(filepath.Join(out, imageName + ".jpeg"))
	if err != nil {
		return errors.New("fail to create file")
	}

	if err := jpeg.Encode(result, img, &jpeg.Options{}); err != nil {
		result.Close()
		return errors.New(err.Error())
	}

	if err := result.Close(); err != nil {
		return errors.New(err.Error())
	}

	return nil
}