package formats

import (
	"encoding/base64"
	"errors"
	"fmt"
	"image/color"
)

type MacOsTerminalFormat struct{}

func (f *MacOsTerminalFormat) FriendlyName() string {
	return "macOS Terminal"
}

func (f *MacOsTerminalFormat) FlagName() string {
	return "macosterminal"
}

func (f *MacOsTerminalFormat) Output() func(colors []color.Color) string {
	return func(colors []color.Color) string {
		const OSXSerializedNSColorTemplate = `<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd"><plist version="1.0"><dict><key>$archiver</key><string>NSKeyedArchiver</string><key>$objects</key><array><string>$null</string><dict><key>$class</key><dict><key>CF$UID</key><integer>2</integer></dict><key>NSColorSpace</key><integer>1</integer><key>NSRGB</key><data>%s</data></dict><dict><key>$classes</key><array><string>NSColor</string><string>NSObject</string></array><key>$classname</key><string>NSColor</string></dict></array><key>$top</key><dict><key>root</key><dict><key>CF$UID</key><integer>1</integer></dict></dict><key>$version</key><integer>100000</integer></dict></plist>`
		OSXColorNames := map[int]string{
			0: "Black",
			1: "Red",
			2: "Green",
			3: "Yellow",
			4: "Blue",
			5: "Magenta",
			6: "Cyan",
			7: "White",
		}

		output := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
		output += "<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">\n"
		output += "<plist version=\"1.0\">\n"
		output += "<dict>\n"
		for i, c := range colors {
			cc := c.(color.NRGBA)
			output += "\t<key>ANSI"
			if i > 7 {
				output += "Bright" + OSXColorNames[i-8]
			} else {
				output += OSXColorNames[i]
			}
			output += "Color</key>\n"
			output += "\t<data>\n"
			rgbColorString := fmt.Sprintf("%.10f %.10f %.10f", float64(cc.R)/255, float64(cc.G)/255, float64(cc.B)/255)
			serializedColor := fmt.Sprintf(OSXSerializedNSColorTemplate, base64.StdEncoding.EncodeToString([]byte(rgbColorString)))
			output += "\t" + base64.StdEncoding.EncodeToString([]byte(serializedColor))
			output += "\n\t</data>\n"
		}

		output += "\t<key>type</key>\n"
		output += "\t<string>Window Settings</string>\n"
		output += "</dict>\n"
		output += "</plist>\n"
		return output
	}
}

func (f *MacOsTerminalFormat) Input() func(filename string) ([]color.Color, error) {
	return func(filename string) ([]color.Color, error) {
		return nil, errors.New("Not yet implemented")
	}
}
