package gummi

import (
	"fmt"
	"os"
	"strings"
)

func prompt() string {
	fmt.Printf("%s> ", userName)
	if !userInput.Scan() {
		fmt.Println()
		fmt.Println("Shutting down.")
		os.Exit(0)
	}
	return userInput.Text()
}

func yesNoPrompt() bool {

	for {
		answer := strings.ToLower(prompt())
		switch answer[0] {
		case 'y':
			return true
		case 'n':
			return false
		default:
			gummiSay("I didn't get that. Can you answer again with yes or no?")
		}
	}
}
