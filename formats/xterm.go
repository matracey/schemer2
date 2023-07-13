package formats

import (
	"encoding/hex"
	"image/color"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type XtermTerminalFormat struct{}

func (f *XtermTerminalFormat) FriendlyName() string {
	return "rxvt/xterm/aterm"
}

func (f *XtermTerminalFormat) FlagName() string {
	return "xterm"
}

func (f *XtermTerminalFormat) Output() func(colors []color.Color) string {
	return func(colors []color.Color) string {
		output := ""
		output += "! Terminal colors"
		output += "\n"
		for i, c := range colors {
			cc := c.(color.NRGBA)
			bytes := []byte{byte(cc.R), byte(cc.G), byte(cc.B)}
			output += "*color"
			output += strconv.Itoa(i)
			output += ": #"
			output += hex.EncodeToString(bytes)
			output += "\n"
		}

		return output
	}
}

func (f *XtermTerminalFormat) Input() func(filename string) ([]color.Color, error) {
	return func(filename string) ([]color.Color, error) {
		config, err := readFile(filename)
		if err != nil {
			return nil, err
		}

		// Split the configuration file into lines
		lines := strings.Split(config, "\n")

		// Remove all spaces from each line
		for i, line := range lines {
			lines[i] = strings.Replace(line, " ", "", -1)
		}

		// Find all lines that contain a color definition
		colorLines := make([]string, 0)
		colorRegex := regexp.MustCompile("[\\*]?[URXvurxterm]*[\\*.]+color[0-9]*")
		for _, line := range lines {
			if len(colorRegex.FindAllString(line, 1)) != 0 {
				colorLines = append(colorLines, line)
			}
		}

		// Sort the color lines by their index
		sort.Slice(colorLines, func(i, j int) bool {
			indexA, _ := strconv.Atoi(strings.TrimSpace(colorLines[i][len(colorLines[i])-1:]))
			indexB, _ := strconv.Atoi(strings.TrimSpace(colorLines[j][len(colorLines[j])-1:]))
			return indexA < indexB
		})

		// Parse each color string and append it to the colors slice
		colors := make([]color.Color, 0)
		for _, line := range colorLines {
			splits := strings.Split(line, ":")
			colorString := splits[len(splits)-1]
			col, err := parseColor(colorString)
			if err != nil {
				return nil, err
			}
			colors = append(colors, col)
		}

		return colors, nil
	}
}
