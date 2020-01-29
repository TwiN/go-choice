package gochoice

import (
	"errors"
	"github.com/gdamore/tcell"
	"strings"
)

var (
	ErrNoChoiceSelected = errors.New("no choice selected")
	ErrNoChoice         = errors.New("no choices to choose from")
	defaultConfig       = Config{
		TextColor:         White.toTcellColor(),
		BackgroundColor:   Black.toTcellColor(),
		SelectedTextColor: White.toTcellColor(),
		SelectedTextBold:  false,
	}
)

// Pick prompts the user to choose an option from a list of choices
func Pick(question string, choicesToPickFrom []string, options ...Option) (string, int, error) {
	config := defaultConfig
	for _, option := range options {
		option(&config)
	}
	screen, err := createScreen()
	if err != nil {
		return "", 0, err
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor))
	return pick(question, choicesToPickFrom, screen, &config)
}

func pick(question string, choicesToPickFrom []string, screen tcell.Screen, config *Config) (string, int, error) {
	if len(choicesToPickFrom) == 0 {
		return "", 0, ErrNoChoice
	}
	var choices []*Choice
	for i, choice := range choicesToPickFrom {
		choices = append(choices, &Choice{Id: i, Value: choice, Selected: i == 0})
	}

	quit := make(chan struct{})
	var selectedChoice = choices[0]
	go func() {
		for {
			go render(screen, question, choices, config, selectedChoice)
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyUp:
					selectedChoice = moveUp(choices, 1)
				case tcell.KeyDown:
					selectedChoice = moveDown(choices, 1)
				case tcell.KeyHome:
					selectedChoice = moveUp(choices, len(choices))
				case tcell.KeyEnd:
					selectedChoice = moveDown(choices, len(choices))
				case tcell.KeyPgUp:
					selectedChoice = moveUp(choices, computePageSize(screen, question))
				case tcell.KeyPgDn:
					selectedChoice = moveDown(choices, computePageSize(screen, question))
				case tcell.KeyEnter, tcell.KeyRight:
					// The current selected choice is already set, so we just quit
					close(quit)
					return
				case tcell.KeyEscape, tcell.KeyCtrlC:
					// No choices were selected, so we'll set selectedChoice to nil and quit
					selectedChoice = nil
					close(quit)
					return
				case tcell.KeyRune:
					switch ev.Rune() {
					case 'k', 'w': // Up
						selectedChoice = moveUp(choices, 1)
					case 'j', 's': // Down
						selectedChoice = moveDown(choices, 1)
					case ' ', 'l', 'd': // Select
						// The current selected choice is already set, so we just quit
						close(quit)
						return
					case 'q': // Quit
						// No choices were selected, so we'll set selectedChoice to nil and quit
						selectedChoice = nil
						close(quit)
						return
					}
				}
			case *tcell.EventResize:
				screen.Sync()
			}
		}
	}()

	<-quit

	if selectedChoice == nil {
		return "", 0, ErrNoChoiceSelected
	}
	return selectedChoice.Value, selectedChoice.Id, nil
}

func computePageSize(screen tcell.Screen, question string) int {
	_, height := screen.Size()
	questionLines := len(strings.Split(question, "\n"))
	if height > questionLines {
		height -= questionLines + 1
	}
	return height
}

func move(choices []*Choice, increment int) *Choice {
	for i, choice := range choices {
		if choice.Selected {
			if i+increment < len(choices) && i+increment > 0 { // Between 0 and last choice
				choices[i].Selected = false
				choices[i+increment].Selected = true
				return choices[i+increment]
			} else if i+increment >= len(choices) { // Higher than last choice
				choices[i].Selected = false
				choices[len(choices)-1].Selected = true
				return choices[len(choices)-1]
			} else if i+increment <= 0 { // Lower than 0
				choices[i].Selected = false
				choices[0].Selected = true
				return choices[0]
			}
			// Choice didn't change, return it
			return choice
		}
	}
	panic("Something went wrong")
}

func moveUp(choices []*Choice, step int) *Choice {
	return move(choices, -step)
}

func moveDown(choices []*Choice, step int) *Choice {
	return move(choices, step)
}
