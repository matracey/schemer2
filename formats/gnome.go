package formats

import (
	"encoding/hex"
	"errors"
	"image/color"
)

type GnomeTerminalFormat struct{}

func (f *GnomeTerminalFormat) FriendlyName() string {
	return "Gnome Terminal (dconf)"
}

func (f *GnomeTerminalFormat) FlagName() string {
	return "gnome-terminal"
}

func (f *GnomeTerminalFormat) Output() func(colors []color.Color) string {
	return func(colors []color.Color) string {
		output := "#!/usr/bin/env bash\npalette=\"["
		for i, c := range colors {
			cc := c.(color.NRGBA)
			bytes := []byte{byte(cc.R), byte(cc.G), byte(cc.B)}
			output += "'"
			output += "#"
			output += hex.EncodeToString(bytes)
			output += "'"
			if i < len(colors)-1 {
				output += ","
			}
		}
		output += "]\""
		output += "\n"
		output += "default=$(dconf read /org/gnome/terminal/legacy/profiles:/default | sed -e \"s/'//g\")"
		output += "\n"
		output += "dconf write /org/gnome/terminal/legacy/profiles:/:$default/palette \"$palette\""
		output += "\n"
		return output
	}
}

func (f *GnomeTerminalFormat) Input() func(filename string) ([]color.Color, error) {
	return func(filename string) ([]color.Color, error) {
		return nil, errors.New("Not yet implemented")
	}
}
