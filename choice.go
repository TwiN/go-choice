package gochoice

import (
	"errors"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"strings"
)

var defaultConfig = Config{
	TextColor:         White,
	BackgroundColor:   Black,
	SelectedTextColor: White,
	SelectedTextBold:  false,
}

func Pick(question string, choicesToPickFrom []string, options ...Option) (string, error) {
	config := defaultConfig
	for _, option := range options {
		option(&config)
	}
	return pick(question, choicesToPickFrom, &config)
}

func pick(question string, choicesToPickFrom []string, config *Config) (string, error) {
	if len(choicesToPickFrom) == 0 {
		return "", errors.New("no choices to choose from")
	}
	var choices []*Choice
	for i, choice := range choicesToPickFrom {
		choices = append(choices, &Choice{Id: i, Value: choice, Selected: i == 0})
	}
	if err := termbox.Init(); err != nil {
		return "", err
	}
	termbox.SetInputMode(termbox.InputEsc)
	defer termbox.Close()
	var selectedChoice = choices[0]
	for {
		render(question, choices, config, selectedChoice)
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

func move(choices []*Choice, increment int) *Choice {
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

func moveUp(choices []*Choice) *Choice {
	return move(choices, -1)
}

func moveDown(choices []*Choice) *Choice {
	return move(choices, 1)
}

func render(question string, options []*Choice, config *Config, selectedChoice *Choice) {
	check(termbox.Clear(config.BackgroundColor.toTermboxAttribute(), config.BackgroundColor.toTermboxAttribute()))
	_, maximumThatCanBeDisplayed := termbox.Size()
	lineNumber := 0
	questionLines := strings.Split(question, "\n")
	for _, line := range questionLines {
		printText(1, lineNumber, line, config.TextColor.toTermboxAttribute(), config.BackgroundColor.toTermboxAttribute())
		lineNumber += 1
	}
	min := selectedChoice.Id + len(questionLines)
	max := maximumThatCanBeDisplayed
	if selectedChoice.Id > max {
		min += 1
		max += 1
	}
	for _, option := range options {
		if option.Id <= (min+1)-maximumThatCanBeDisplayed && !(option.Id > (min+1)-maximumThatCanBeDisplayed) {
			continue
		}
		if option.Selected {
			selectedTextAttribute := config.SelectedTextColor.toTermboxAttribute()
			if config.SelectedTextBold {
				selectedTextAttribute |= termbox.AttrBold
			}
			printText(1, lineNumber, "> "+option.Value, selectedTextAttribute, config.BackgroundColor.toTermboxAttribute())
		} else {
			printText(3, lineNumber, option.Value, config.TextColor.toTermboxAttribute(), config.BackgroundColor.toTermboxAttribute())
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
