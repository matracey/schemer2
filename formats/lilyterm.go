package formats

import (
	"encoding/hex"
	"image/color"
	"strconv"
	"strings"
)

type LilytermTerminalFormat struct{}

func (f *LilytermTerminalFormat) FriendlyName() string {
	return "LilyTerm"
}

func (f *LilytermTerminalFormat) FlagName() string {
	return "lilyterm"
}

func (f *LilytermTerminalFormat) Output() func(colors []color.Color) string {
	return func(colors []color.Color) string {
		output := ""
		for i, c := range colors {
			cc := c.(color.NRGBA)
			bytes := []byte{byte(cc.R), byte(cc.G), byte(cc.B)}
			output += "Color"
			output += strconv.Itoa(i)
			output += " = "
			output += "#"
			output += hex.EncodeToString(bytes)
			output += "\n"
		}
		return output
	}
}

func (f *LilytermTerminalFormat) Input() func(filename string) ([]color.Color, error) {
	return func(filename string) ([]color.Color, error) {
		colors := make([]color.Color, 0)
		config, err := readFile(filename)
		if err != nil {
			return nil, err
		}
		lines := strings.Split(config, "\n")
		for i, line := range lines {
			lines[i] = strings.Replace(line, " ", "", -1)
		}
		for i := 0; i < 16; i++ {
			for _, line := range lines {
				prefix := "Color"
				prefix += strconv.Itoa(i)
				prefix += "="
				if strings.HasPrefix(line, prefix) {
					hexString := strings.TrimPrefix(line, prefix)

					col, err := parseColor(hexString)
					if err != nil {
						return nil, err
					}

					colors = append(colors, col)

				}
			}
		}

		return colors, nil
	}
}
