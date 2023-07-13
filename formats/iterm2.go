package formats

import (
	"errors"
	"image/color"
	"strconv"
)

type ITerm2TerminalFormat struct{}

func (f *ITerm2TerminalFormat) FriendlyName() string {
	return "iTerm2"
}

func (f *ITerm2TerminalFormat) FlagName() string {
	return "iterm2"
}

func (f *ITerm2TerminalFormat) Output() func(colors []color.Color) string {
	return func(colors []color.Color) string {
		output := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
		output += "<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">\n"
		output += "<plist version=\"1.0\">\n"
		output += "<dict>\n"
		for i, c := range colors {
			cc := c.(color.NRGBA)
			output += "\t<key>Ansi "
			output += strconv.Itoa(i)
			output += " Color</key>\n"
			output += "\t<dict>\n"
			output += "\t\t<key>Blue Component</key>\n"
			output += "\t\t<real>"
			output += strconv.FormatFloat(float64(cc.B)/255, 'f', 17, 64)
			output += "</real>\n"
			output += "\t\t<key>Green Component</key>\n"
			output += "\t\t<real>"
			output += strconv.FormatFloat(float64(cc.G)/255, 'f', 17, 64)
			output += "</real>\n"
			output += "\t\t<key>Red Component</key>\n"
			output += "\t\t<real>"
			output += strconv.FormatFloat(float64(cc.R)/255, 'f', 17, 64)
			output += "</real>\n"
			output += "\t</dict>\n"
		}
		output += "</dict>\n"
		output += "</plist>\n"
		return output
	}
}

func (f *ITerm2TerminalFormat) Input() func(filename string) ([]color.Color, error) {
	return func(filename string) ([]color.Color, error) {
		return nil, errors.New("Not yet implemented")
	}
}
