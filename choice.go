package gochoice

import (
	"errors"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"strings"
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

func Pick(question string, choicesToPickFrom []string) (string, error) {
	if len(choicesToPickFrom) == 0 {
		return "", errors.New("no choices to choose from")
	}
	var choices []Choice
	for i, choice := range choicesToPickFrom {
		choices = append(choices, Choice{Value: choice, Selected: i == 0})
	}
	if err := termbox.Init(); err != nil {
		return "", err
	}
	defer termbox.Close()
	var selectedChoice = choices[0]
	for {
		render(question, choices)
		switch ev := termbox.PollEvent(); ev.Ch {
		case 0:
			switch ev.Key {
			case termbox.KeyArrowUp:
				selectedChoice = moveUp(choices)
			case termbox.KeyArrowDown:
				selectedChoice = moveDown(choices)
			case termbox.KeyEnter, termbox.KeySpace:
				return selectedChoice.Value, nil
			case termbox.KeyEsc:
				return "", errors.New("aborted")
			default:
			}
		case 'q':
			return "", errors.New("aborted")
		case 'k', 'w': // up
			selectedChoice = moveUp(choices)
		case 'j', 's': // down
			selectedChoice = moveDown(choices)
		case 'l', 'd': // right
			return selectedChoice.Value, nil
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

func moveUp(choices []Choice) Choice {
	return move(choices, -1)
}

func moveDown(choices []Choice) Choice {
	return move(choices, 1)
}

func render(question string, options []Choice) {
	check(termbox.Clear(bgColor, bgColor))
	lineNumber := 0
	for _, line := range strings.Split(question, "\n") {
		printText(1, lineNumber, line, fgColor, bgColor)
		lineNumber += 1
	}

	for _, option := range options {
		if option.Selected {
			printText(1, lineNumber, "> "+option.Value, fgColor, bgColor)
		} else {
			printText(3, lineNumber, option.Value, fgColor, bgColor)
		}
		lineNumber += 1
	}
	check(termbox.Flush())
}

func printText(x, y int, text string, fg, bg termbox.Attribute) {
	for _, character := range text {
		termbox.SetCell(x, y, character, fg, bg)
		x += runewidth.RuneWidth(character)
	}
}
