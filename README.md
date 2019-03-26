# go-choice

A very simple library for interactively selecting an option on a terminal 


## Usage

```
package main

import (
	"github.com/TwinProduction/go-choice"
)

func main() {
	choice := gochoice.Pick(
		"What do you want to do?",
		[]string{
			"Connect to the production environment",
			"Connect to the test environment",
			"Update",
		})
	println("You have selected: " + choice)
}
```

![example](assets/example.gif)