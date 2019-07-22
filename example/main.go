package main

import (
	"github.com/TwinProduction/go-choice"
)

func main() {
	choice, err := gochoice.Pick(
		"What do you want to do?",
		[]string{
			"Connect to the production environment",
			"Connect to the test environment",
			"Update",
		})
	if err != nil {
		println("You didn't select anything!")
	} else {
		println("You have selected: " + choice)
	}
}
