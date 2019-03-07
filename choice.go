package gochoice

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Choice struct {
	Value    string
	Selected bool
	//Disabled bool
}

const (
	fgColor = termbox.ColorWhite
	bgColor = termbox.ColorBlack
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Pick(question string, choicesToPickFrom []string) string {
	if len(choicesToPickFrom) == 0 {
		panic("No choices to choose from")
	}
	var choices []Choice
	for i, choice := range choicesToPickFrom {
		choices = append(choices, Choice{Value: choice, Selected: i == 0})
	}
	check(termbox.Init())
	defer termbox.Close()
	var selectedChoice = choices[0]
	for {
		render(question, choices)
		switch ev := termbox.PollEvent(); ev.Key {
		case termbox.KeyArrowDown:
			selectedChoice = move(choices, 1)
		case termbox.KeyArrowUp:
			selectedChoice = move(choices, -1)
		case termbox.KeyEnter:
			return selectedChoice.Value
		case termbox.KeyEsc:
			panic("Aborted")
		default:
		}
	}
}

func move(choices []Choice, increment int) Choice {
	for i, choice := range choices {
		if choice.Selected {
			if i+increment < len(choices) && i+increment >= 0 {
				choices[i].Selected = false
				choices[i+increment].Selected = true
				return choices[i+increment]
			}
			// Choice didn't change, return it
			return choice
		}
	}
	panic("Something went wrong")
}

func render(question string, options []Choice) {
	check(termbox.Clear(bgColor, bgColor))
	printText(1, 0, question, fgColor, bgColor)
	for i, option := range options {
		lineNumber := i + 1
		if option.Selected {
			printText(1, lineNumber, "> "+option.Value, fgColor, bgColor)
		} else {
			printText(3, lineNumber, option.Value, fgColor, bgColor)
		}
	}
	check(termbox.Flush())
}

func printText(x, y int, text string, fg, bg termbox.Attribute) {
	for _, character := range text {
		termbox.SetCell(x, y, character, fg, bg)
		x += runewidth.RuneWidth(character)
	}
}
