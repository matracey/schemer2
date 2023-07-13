package formats

import (
	"encoding/hex"
	"errors"
	"image/color"
)

type ColorsFormat struct{}

func (f *ColorsFormat) FriendlyName() string {
	return "Colors in Plain Text"
}

func (f *ColorsFormat) FlagName() string {
	return "colors"
}

func (f *ColorsFormat) Output() func(colors []color.Color) string {
	return func(colors []color.Color) string {
		output := ""
		for i, c := range colors {
			cc := c.(color.NRGBA)
			bytes := []byte{byte(cc.R), byte(cc.G), byte(cc.B)}
			output += "#"
			output += hex.EncodeToString(bytes)
			output += "\n"
		}
		return output
	}
}

func (f *ColorsFormat) Input() func(filename string) ([]color.Color, error) {
	return func(filename string) ([]color.Color, error) {
		return nil, errors.New("Not yet implemented")
	}
}
