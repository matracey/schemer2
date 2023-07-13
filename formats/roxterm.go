package formats

type RoxtermTerminalFormat struct{}

func (f *RoxtermTerminalFormat) FriendlyName() string {
	return "ROXTerm"
}

func (f *RoxtermTerminalFormat) FlagName() string {
	return "roxterm"
}

func (f *RoxtermTerminalFormat) Output() func(colors []color.Color) string { return func (colors []color.Color) string {
	output := "[roxterm colour scheme]\n"
	output += "pallete
size=16\n"

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

func (f *RoxtermTerminalFormat) Input() (func(filename string) ([]color.Color, error)) {}
