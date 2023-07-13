package formats

import (
	"errors"
	"image/color"
	"strconv"
)

type KonsoleTerminalFormat struct{}

func (f *KonsoleTerminalFormat) FriendlyName() string {
	return "Konsole"
}

func (f *KonsoleTerminalFormat) FlagName() string {
	return "konsole"
}

func (f *KonsoleTerminalFormat) Output() func(colors []color.Color) string {
	return func(colors []color.Color) string {
		output := ""
		for i, c := range colors {
			cc := c.(color.NRGBA)
			output += "[Color"
			if i > 7 {
				output += strconv.Itoa(i - 8)
				output += "Intense"
			} else {
				output += strconv.Itoa(i)
			}
			output += "]\n"
			output += "Color="
			output += strconv.Itoa(int(cc.R)) + ","
			output += strconv.Itoa(int(cc.G)) + ","
			output += strconv.Itoa(int(cc.B)) + "\n"
			output += "Transparency=false\n\n"
		}

		return output
	}
}

func (f *KonsoleTerminalFormat) Input() func(filename string) ([]color.Color, error) {
	return func(filename string) ([]color.Color, error) {
		return nil, errors.New("Not yet implemented")
	}
}
