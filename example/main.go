package main

import (
	"github.com/TwinProduction/go-choice"
)

func main() {
	choice := gochoice.Pick(
		"What do you want to do?\nPick one option below",
		[]string{
			"Connect to the production environment",
			"Connect to the test environment",
			"Update",
		})
	println("You have selected: " + choice)
}
