package formats

import (
	"errors"
	"image/color"
)

type ImgFormat struct{}

func (f *ImgFormat) FriendlyName() string {
	return "Image"
}

func (f *ImgFormat) FlagName() string {
	return "img"
}

func (f *ImgFormat) Output() func(colors []color.Color) string {
	return func(colors []color.Color) string {
		return "Not yet implemented"
	}
}

func (f *ImgFormat) Input() func(filename string) ([]color.Color, error) {
	return func(filename string) ([]color.Color, error) {
		return nil, errors.New("Not yet implemented")
	}
}
