package main

import (
	"flag"
	"strings"
)

type Colors string

var colors = flag.String("color", "color name", "display colorized output")

const (
	Black                   Colors = "\u001b[30m"
	Red                            = "\u001b[31m"
	Green                          = "\u001b[32m"
	Yellow                         = "\u001b[38;2;255;255;0m"
	Orange                         = "\u001b[33m"
	Blue                           = "\u001b[34m"
	Magenta                        = "\u001b[35m"
	Cyan                           = "\u001b[36m"
	White                          = "\u001b[37m"
	BrightBlack                    = "\u001b[30;1m"
	BrightRed                      = "\u001b[31;1m"
	BrightGreen                    = "\u001b[32;1m"
	BrightYellow                   = "\u001b[33;1m"
	BrightBlue                     = "\u001b[34;1m"
	BrightMagenta                  = "\u001b[35;1m"
	BrightCyan                     = "\u001b[36;1m"
	BrightWhite                    = "\u001b[37;1m"
	BackgroundBlack                = "\u001b[40m"
	BackgroundRed                  = "\u001b[41m"
	BackgroundGreen                = "\u001b[42m"
	BackgroundYellow               = "\u001b[43m"
	BackgroundBlue                 = "\u001b[44m"
	BackgroundMagenta              = "\u001b[45m"
	BackgroundCyan                 = "\u001b[46m"
	BackgroundWhite                = "\u001b[47m"
	BackgroundBrightBlack          = "\u001b[40;1m"
	BackgroundBrightRed            = "\u001b[41;1m"
	BackgroundBrightGreen          = "\u001b[42;1m"
	BackgroundBrightYellow         = "\u001b[43;1m"
	BackgroundBrightBlue           = "\u001b[44;1m"
	BackgroundBrightMagenta        = "\u001b[45;1m"
	BackgroundBrightCyan           = "\u001b[46;1m"
	BackgroundBrightWhite          = "\u001b[47;1m"
	Reset                          = "\u001b[0m"
)

// Map of color names to ANSI color codes
var colorMap = map[string]Colors{
	"black":                   Black,
	"red":                     Red,
	"green":                   Green,
	"yellow":                  Yellow,
	"orange":                  Orange,
	"blue":                    Blue,
	"magenta":                 Magenta,
	"cyan":                    Cyan,
	"white":                   White,
	"brightBlack":             BrightBlack,
	"brightRed":               BrightRed,
	"brightGreen":             BrightGreen,
	"brightYellow":            BrightYellow,
	"brightBlue":              BrightBlue,
	"brightMagenta":           BrightMagenta,
	"brightCyan":              BrightCyan,
	"brightWhite":             BrightWhite,
	"backgroundBlack":         BackgroundBlack,
	"backgroundRed":           BackgroundRed,
	"backgroundGreen":         BackgroundGreen,
	"backgroundYellow":        BackgroundYellow,
	"backgroundBlue":          BackgroundBlue,
	"backgroundMagenta":       BackgroundMagenta,
	"backgroundCyan":          BackgroundCyan,
	"backgroundWhite":         BackgroundWhite,
	"backgroundBrightBlack":   BackgroundBrightBlack,
	"backgroundBrightRed":     BackgroundBrightRed,
	"backgroundBrightGreen":   BackgroundBrightGreen,
	"backgroundBrightYellow":  BackgroundBrightYellow,
	"backgroundBrightBlue":    BackgroundBrightBlue,
	"backgroundBrightMagenta": BackgroundBrightMagenta,
	"backgroundBrightCyan":    BackgroundBrightCyan,
	"backgroundBrightWhite":   BackgroundBrightWhite,
	"reset":                   Reset,
}

// get color
func getANSIColor(message string) []string {
	color, ok := colorMap[*colors]
	if !ok {
		color = Reset
		// fmt.Println("\nColor not available.\nPlease select one of the available colors from README.MD")
		// os.Exit(0)

	}

	return []string{colorize(color, message)}
}

// color the message
func colorize(color Colors, message string) string {
	message = string(color) + message + string(Reset)
	return message
}

// color the output and join it together to string
func color(ascii string) string {
	color := getANSIColor(ascii)
	coloredString := strings.Join(color, "")
	return coloredString
}
