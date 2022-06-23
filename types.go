package gochoice

import (
	"github.com/gdamore/tcell/v2"
)

type choice struct {
	ID       int
	Value    string
	Selected bool

	hidden bool
}

type Config struct {
	TextColor         tcell.Color
	BackgroundColor   tcell.Color
	SelectedTextColor tcell.Color
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
	Orange
	Gold
	Gray
	Grey
	Fuchsia
	Brown
	Pink
	Purple
	Crimson
	DarkRed
	DarkBlue
	DarkGray
	DarkGrey
	LightBlue
	LightGray
	LightGrey
	White
)

func (c Color) toTcellColor() tcell.Color {
	switch c {
	case Black:
		return tcell.ColorBlack
	case Red:
		return tcell.ColorRed
	case Green:
		return tcell.ColorGreen
	case Yellow:
		return tcell.ColorYellow
	case Blue:
		return tcell.ColorBlue
	case Magenta:
		return tcell.ColorDarkMagenta
	case Cyan:
		return tcell.ColorLightCyan
	case Orange:
		return tcell.ColorOrange
	case Gold:
		return tcell.ColorGold
	case Gray, Grey:
		return tcell.ColorGray
	case Fuchsia:
		return tcell.ColorFuchsia
	case Brown:
		return tcell.ColorBrown
	case Pink:
		return tcell.ColorPink
	case Purple:
		return tcell.ColorPurple
	case Crimson:
		return tcell.ColorCrimson
	case DarkRed:
		return tcell.ColorDarkRed
	case DarkBlue:
		return tcell.ColorDarkBlue
	case DarkGray, DarkGrey:
		return tcell.ColorDarkGray
	case LightBlue:
		return tcell.ColorLightBlue
	case LightGray, LightGrey:
		return tcell.ColorLightGray
	case White:
		fallthrough
	default:
		return tcell.ColorWhite
	}
}

type Option func(config *Config)

func OptionTextColor(color Color) func(config *Config) {
	return func(config *Config) {
		config.TextColor = color.toTcellColor()
	}
}

func OptionBackgroundColor(color Color) func(config *Config) {
	return func(config *Config) {
		config.BackgroundColor = color.toTcellColor()
	}
}

func OptionSelectedTextColor(color Color) func(config *Config) {
	return func(config *Config) {
		config.SelectedTextColor = color.toTcellColor()
	}
}

func OptionSelectedTextBold() func(config *Config) {
	return func(config *Config) {
		config.SelectedTextBold = true
	}
}
