package formats

import (
	"encoding/hex"
	"errors"
	"image/color"
	"strconv"
)

type UrxvtTerminalFormat struct{}

func (f *UrxvtTerminalFormat) FriendlyName() string {
	return "urxvt"
}

func (f *UrxvtTerminalFormat) FlagName() string {
	return "urxvt"
}

func (f *UrxvtTerminalFormat) Output() func(colors []color.Color) string {
	return func(colors []color.Color) string {
		output := ""
		for i, c := range colors {
			cc := c.(color.NRGBA)
			bytes := []byte{byte(cc.R), byte(cc.G), byte(cc.B)}
			output += "URxvt*color"
			output += strconv.Itoa(i)
			output += ": "
			output += "#"
			output += hex.EncodeToString(bytes)
			output += "\n"
		}
		return output
	}
}

func (f *UrxvtTerminalFormat) Input() func(filename string) ([]color.Color, error) {
	return func(filename string) ([]color.Color, error) {
		return nil, errors.New("Not yet implemented")
	}
}
