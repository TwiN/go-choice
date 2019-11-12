package gochoice

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
	"strconv"
	"strings"
)

func createScreen() (tcell.Screen, error) {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("failed to create new screen: %v", err)
	}
	if err := screen.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize screen: %v", err)
	}
	return screen, nil
}

// Renders content on the screen
func render(screen tcell.Screen, question string, options []*Choice, config *Config, selectedChoice *Choice) {
	_, screenHeight := screen.Size()
	lineNumber := 0
	// Display question
	questionLines := strings.Split(question, "\n")
	for _, questionLine := range questionLines {
		printText(screen, 0, lineNumber, fmt.Sprintf(" %s", questionLine), config.TextColor, config.BackgroundColor, config.SelectedTextBold)
		lineNumber += 1
	}
	// Display all options that can fit in the screen
	min := selectedChoice.Id + len(questionLines)
	for i, option := range options {
		if option.Id <= (min+2)-screenHeight && !(option.Id > (min+2)-screenHeight) && i-screenHeight < min {
			continue
		}
		if option.Selected {
			printText(screen, 0, lineNumber, fmt.Sprintf(" > %s", option.Value), config.SelectedTextColor, config.BackgroundColor, config.SelectedTextBold)
		} else {
			printText(screen, 0, lineNumber, fmt.Sprintf("   %s", option.Value), config.TextColor, config.BackgroundColor, config.SelectedTextBold)
		}
		lineNumber += 1
	}
	// HACK: Instead of using screen.Clear(), draw over the existing text
	for i := lineNumber; i < screenHeight; i++ {
		printText(screen, 1, i, "", config.TextColor, config.BackgroundColor, config.SelectedTextBold)
	}
	screen.Show()
}

// Prints text a screen
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
