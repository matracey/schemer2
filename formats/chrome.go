package formats

import (
	"encoding/hex"
	"errors"
	"image/color"
	"strconv"
)

type ChromeTerminalFormat struct{}

func (f *ChromeTerminalFormat) FriendlyName() string {
	return "Chrome Shell"
}

func (f *ChromeTerminalFormat) FlagName() string {
	return "chrome"
}

func (f *ChromeTerminalFormat) Output() func(colors []color.Color) string {
	return func(colors []color.Color) string {
		output := "{"
		for i, c := range colors {
			cc := c.(color.NRGBA)
			bytes := []byte{byte(cc.R), byte(cc.G), byte(cc.B)}
			output += " \""
			output += strconv.Itoa(i)
			output += "\""
			output += ": "
			output += " \""
			output += "#"
			output += hex.EncodeToString(bytes)
			output += "\" "
			if i != len(colors)-1 {
				output += ", "
			}
		}
		output += "}\n"
		return output
	}
}

func (f *ChromeTerminalFormat) Input() func(filename string) ([]color.Color, error) {
	return func(filename string) ([]color.Color, error) {
		return nil, errors.New("Not yet implemented")
	}
}
