package gochoice

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

func createScreen() (tcell.Screen, error) {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("failed to create new screen: %w", err)
	}
	if err := screen.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize screen: %w", err)
	}
	return screen, nil
}

// render the question, options and the selected choice with the given configuration
func render(screen tcell.Screen, question string, options []*choice, config *Config, searchQuery string) {
	_, screenHeight := screen.Size()
	lineNumber := 0
	// Display question
	questionLines := strings.Split(question, "\n")
	for _, questionLine := range questionLines {
		printText(screen, 0, lineNumber, " "+questionLine, config.TextColor, config.BackgroundColor, config.SelectedTextBold)
		lineNumber++
	}
	selectedChoiceIndex := 0
	numberOfOptionsNotHidden := 0
	for _, option := range options {
		if len(searchQuery) > 0 && !strings.Contains(strings.ToLower(option.Value), strings.ToLower(searchQuery)) {
			option.hidden = true
		} else {
			option.hidden = false
			if option.Selected {
				selectedChoiceIndex = numberOfOptionsNotHidden
			}
			numberOfOptionsNotHidden++
		}
	}
	// Display all options that can fit in the screen
	min := selectedChoiceIndex + len(questionLines)
	visibleOptionIndex := 0
	for _, option := range options {
		if option.hidden {
			continue
		}
		visibleOptionIndex++
		if visibleOptionIndex <= (min+2)-screenHeight && !(visibleOptionIndex > (min+2)-screenHeight) && visibleOptionIndex-screenHeight < min {
			continue
		}
		if option.Selected {
			printText(screen, 0, lineNumber, " > "+option.Value, config.SelectedTextColor, config.BackgroundColor, config.SelectedTextBold)
		} else {
			printText(screen, 0, lineNumber, "   "+option.Value, config.TextColor, config.BackgroundColor, config.SelectedTextBold)
		}
		lineNumber++
	}
	if numberOfOptionsNotHidden == 0 {
		printText(screen, 1, lineNumber, " ! There are no choices matching your search query", config.TextColor, config.BackgroundColor, config.SelectedTextBold)
		lineNumber++
	}
	// HACK: Instead of using screen.Clear(), draw over the existing text
	for i := lineNumber; i < screenHeight; i++ {
		printText(screen, 1, i, "", config.TextColor, config.BackgroundColor, config.SelectedTextBold)
	}
	printText(screen, 1, screenHeight-1, "Search: "+searchQuery+"_", config.TextColor, config.BackgroundColor, config.SelectedTextBold)
	screen.Show()
}

// printText prints text on the given screen
func printText(screen tcell.Screen, x, y int, text string, fg, bg tcell.Color, bold bool) {
	// Overwrite all existing characters on the line with the new text
	width, _ := screen.Size()
	textWithSpaces := fmt.Sprintf("%-"+strconv.Itoa(width)+"s", text)
	// Write all characters on the screen
	for _, character := range textWithSpaces {
		screen.SetCell(x, y, tcell.StyleDefault.Background(bg).Foreground(fg).Bold(bold), character)
		x += runewidth.RuneWidth(character)
	}
}
