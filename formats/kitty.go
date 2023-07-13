package formats

import (
	"encoding/hex"
	"image/color"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type KittyTerminalFormat struct{}

func (f *KittyTerminalFormat) FriendlyName() string {
	return "Kitty Terminal"
}

func (f *KittyTerminalFormat) FlagName() string {
	return "kitty"
}

func (f *KittyTerminalFormat) Output() func(colors []color.Color) string {
	return func(colors []color.Color) string {
		output := ""
		for i, c := range colors {
			cc := c.(color.NRGBA)
			bytes := []byte{byte(cc.R), byte(cc.G), byte(cc.B)}
			output += "color"
			output += strconv.Itoa(i)
			output += "\t#"
			output += hex.EncodeToString(bytes)
			output += "\n"
		}

		return output
	}
}

func (f *KittyTerminalFormat) Input() func(filename string) ([]color.Color, error) {
	return func(filename string) ([]color.Color, error) {
		config, err := readFile(filename)
		if err != nil {
			return nil, err
		}

		lines := strings.Split(config, "\n")

		colorLines := make([]string, 0)
		colorRegex := regexp.MustCompile("^color[0-9]*")
		for _, line := range lines {
			if len(colorRegex.FindAllString(line, 1)) != 0 {
				colorLines = append(colorLines, strings.Replace(line, "color", "", 1))
			}
		}

		sort.Slice(colorLines, func(i, j int) bool {
			indexA, _ := strconv.Atoi(strings.TrimSpace(colorLines[i][:2]))
			indexB, _ := strconv.Atoi(strings.TrimSpace(colorLines[j][:2]))
			return indexA < indexB
		})

		colors := make([]color.Color, 0)
		for _, line := range colorLines {
			splits := strings.Fields(line)
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
