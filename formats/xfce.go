package formats

import (
	"encoding/hex"
	"errors"
	"image/color"
	"strings"
)

type XfceTerminalFormat struct{}

func (f *XfceTerminalFormat) FriendlyName() string {
	return "XFCE4Terminal"
}

func (f *XfceTerminalFormat) FlagName() string {
	return "xfce"
}

func (f *XfceTerminalFormat) Output() func(colors []color.Color) string {
	return func(colors []color.Color) string {
		output := ""
		output += "ColorPalette="
		for _, c := range colors {
			bytes := []byte{byte(c.(color.NRGBA).R), byte(c.(color.NRGBA).R), byte(c.(color.NRGBA).G), byte(c.(color.NRGBA).G), byte(c.(color.NRGBA).B), byte(c.(color.NRGBA).B)}
			output += "#"
			output += hex.EncodeToString(bytes)
			output += "	"
		}
		output += "\n"

		return output
	}
}

func (f *XfceTerminalFormat) Input() func(filename string) ([]color.Color, error) {
	return func(filename string) ([]color.Color, error) {
		config, err := readFile(filename)
		if err != nil {
			return nil, err
		}
		lines := strings.Split(config, "\n")
		for i, line := range lines {
			lines[i] = strings.Replace(line, " ", "", -1)
		}
		colorPalette := ""
		for _, line := range lines {
			if strings.HasPrefix(line, "ColorPalette") {
				colorPalette = line
			}
		}
		if colorPalette == "" {
			return nil, errors.New("ColorPalette not found in XFCE4 Terminal input")
		}
		colorPalette = strings.TrimPrefix(colorPalette, "ColorPalette=")
		colorPalette = strings.TrimRight(colorPalette, "; ")
		colorStrings := strings.Split(colorPalette, ";")

		colors := make([]color.Color, 0)

		for _, colorString := range colorStrings {
			col, err := parseColor(colorString)
			if err != nil {
				return nil, err
			}
			colors = append(colors, col)

		}
		return colors, nil
	}
}
