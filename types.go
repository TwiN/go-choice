package gochoice

import "github.com/nsf/termbox-go"

type Choice struct {
	Id       int
	Value    string
	Selected bool
}

type Config struct {
	TextColor         Color
	BackgroundColor   Color
	SelectedTextColor Color
	SelectedTextBold  bool
}

type Color int

const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func (c Color) toTermboxAttribute() termbox.Attribute {
	switch c {
	case Black:
		return termbox.ColorBlack
	case Red:
		return termbox.ColorRed
	case Green:
		return termbox.ColorGreen
	case Yellow:
		return termbox.ColorYellow
	case Blue:
		return termbox.ColorBlue
	case Magenta:
		return termbox.ColorMagenta
	case Cyan:
		return termbox.ColorCyan
	default:
		return termbox.ColorWhite
	}
}

type Option func(config *Config)

func OptionTextColor(color Color) func(config *Config) {
	return func(config *Config) {
		config.TextColor = color
	}
}

func OptionBackgroundColor(color Color) func(config *Config) {
	return func(config *Config) {
		config.BackgroundColor = color
	}
}

func OptionSelectedTextColor(color Color) func(config *Config) {
	return func(config *Config) {
		config.SelectedTextColor = color
	}
}

func OptionSelectedTextBold() func(config *Config) {
	return func(config *Config) {
		config.SelectedTextBold = true
	}
}
