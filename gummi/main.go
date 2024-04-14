package gummi

import (
	"bufio"
	"fmt"
	"os"
	"os/user"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/validate"
)

func Setup() {
	setupUserName()
	setupUserInput()
}

func Intro(f *validate.Functions) {
	printHello()
	printIssue(f)
}

func Fix(f *validate.Functions) {
	if f.Sammi != nil {
		fixSammi(f)
	}
	config.Config.Save()
	if f.Gumroad != nil {
		fixGumroad(f)
	}
	config.Config.Save()
	if f.Valid() {
		printSucceeded()
	}

	err := config.Config.Save()
	if err != nil {
		fmt.Println(err)
	}
}

func setupUserName() {
	u, err := user.Current()
	if err != nil {
		userName = "#####"
	} else {
		userName = u.Username
	}
}

func setupUserInput() {
	userInput = bufio.NewScanner(os.Stdin)
	userInput.Split(bufio.ScanLines)
}
