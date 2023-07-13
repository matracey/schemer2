package formats

import (
	"encoding/hex"
	"image/color"
	"strconv"
	"strings"
)

type TermiteTerminalFormat struct{}

func (f *TermiteTerminalFormat) FriendlyName() string {
	return "Termite"
}

func (f *TermiteTerminalFormat) FlagName() string {
	return "termite"
}

func (f *TermiteTerminalFormat) Output() func(colors []color.Color) string {
	return func(colors []color.Color) string {
		output := ""
		for i, c := range colors {
			cc := c.(color.NRGBA)
			bytes := []byte{byte(cc.R), byte(cc.G), byte(cc.B)}
			output += "color"
			output += strconv.Itoa(i)
			output += " = "
			output += "#"
			output += hex.EncodeToString(bytes)
			output += "\n"
		}
		return output
	}
}

func (f *TermiteTerminalFormat) Input() func(filename string) ([]color.Color, error) {
	return func(filename string) ([]color.Color, error) {
		colors := make([]color.Color, 0)
		config, err := readFile(filename)
		if err != nil {
			return nil, err
		}
		lines := strings.Split(config, "\n")
		for i, l := range lines {
			lines[i] = strings.Replace(l, " ", "", -1)
		}
		for i := 0; i < 16; i++ {
			for _, l := range lines {
				prefix := "color"
				prefix += strconv.Itoa(i)
				prefix += "="
				if strings.HasPrefix(l, prefix) {
					hexstring := strings.TrimPrefix(l, prefix)

					col, err := parseColor(hexstring)
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
