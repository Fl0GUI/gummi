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

func Intro(f *validate.Functions, config *config.Configuration) {
	printHello()
	printOptions(config)
	validate.UpdateValidation(f, config)
	printIssue(f)
}

func Fix(f *validate.Functions, config *config.Configuration) {
	if f.Sammi != nil {
		fixSammi(f, config)
	}
	config.Save()

	if f.Gumroad != nil {
		fixGumroad(f, config)
	}
	config.Save()

	if config.FourthWallConfig.Active && len(config.FourthWallConfig.AccessToken) == 0 {
		fixFourthWall(f, config)
	}
	config.Save()

	if config.ThroneConfig.Active && len(config.ThroneConfig.CreatorId) == 0 {
		fixThrone(f, config)
	}
	config.Save()

	if f.Valid() {
		printSucceeded()
	}

	err := config.Save()
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
