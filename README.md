# go-choice

A very simple library for interactively selecting an option on a terminal 


## Usage

```go
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
```

![example](assets/example.gif)

Alternatively, you can customize the experience further by using `PickWithConfig` instead of `Pick`:

```go
package main

import (
	"github.com/TwinProduction/go-choice"
)

func main() {
	choice, err := gochoice.PickWithConfig(
		"What do you want to do?\nYour question can also span multiple lines",
		[]string{
			"Connect to the production environment",
			"Connect to the test environment",
			"Update",
		},
		&gochoice.Config{
			BackgroundColor:   gochoice.Black,
			TextColor:         gochoice.White,
			SelectedTextColor: gochoice.Red,
			SelectedTextBold:  true,
		})
	if err != nil {
		println("You didn't select anything!")
	} else {
		println("You have selected: " + choice)
	}
}
```
