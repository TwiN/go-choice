package main

import (
	"fmt"
	"github.com/TwinProduction/go-choice"
)

func main() {
	choice, index, err := gochoice.Pick(
		"What do you want to do?\nPick:",
		[]string{
			"Connect to the production environment",
			"Connect to the test environment",
			"Update",
		}, gochoice.OptionBackgroundColor(gochoice.Black), gochoice.OptionSelectedTextColor(gochoice.Red))
	if err != nil {
		fmt.Println("You didn't select anything!")
	} else {
		fmt.Printf("You have selected: '%s', which is the index %d\n", choice, index)
	}
}
