package formats

import (
	"encoding/hex"
	"errors"
	"image/color"
	"strings"
)

type TerminatorTerminalFormat struct{}

func (f *TerminatorTerminalFormat) FriendlyName() string {
	return "Terminator"
}

func (f *TerminatorTerminalFormat) FlagName() string {
	return "terminator"
}

func (f *TerminatorTerminalFormat) Output() func(colors []color.Color) string {
	return func(colors []color.Color) string {
		output := "palette = \""
		for i, c := range colors {
			cc := c.(color.NRGBA)
			bytes := []byte{byte(cc.R), byte(cc.G), byte(cc.B)}
			if i < len(colors)-1 {
				output += "#"
				output += hex.EncodeToString(bytes)
				output += ":"
			} else if i == len(colors)-1 {
				output += "#"
				output += hex.EncodeToString(bytes)
				output += "\"\n"
			}
		}
		return output
	}
}

func (f *TerminatorTerminalFormat) Input() func(filename string) ([]color.Color, error) {
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
			if strings.HasPrefix(line, "palette") {
				colorPalette = line
			}
		}

		if colorPalette == "" {
			return nil, errors.New("ColorPalette not found in Terminator input")
		}

		colorPalette = strings.TrimPrefix(colorPalette, "palette=\"")
		colorPalette = strings.TrimSuffix(colorPalette, "\"")

		colorStrings := strings.Split(colorPalette, ":")

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
