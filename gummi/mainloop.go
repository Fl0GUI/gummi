package gummi

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"

	"j322.ica/gumroad-sammi/config"
)

func Repl() *config.Config {
	c := config.NewConfig()
	configGood := &functions{false, false, false, false, false}

	setupUserName()
	setupUserInput()

	if config.FileExists() {
		conf, err := config.Load()
		if err != nil {
			log.Printf("Could not load the config correctly: %s\n", err)
		} else {
			c = conf
		}
	} else {
		log.Println("No config file.")
	}
	testConfig(c, configGood)

	if configGood.valid() {
		log.Println("Loaded config is valid, skipping Gummi agent")
		return c
	}

	printHello()
	printIssue(configGood)
	if !configGood.validSammi() {
		fixSammi(c, configGood)
	}
	if !configGood.validGumroad() {
		fixGumroad(c, configGood)
	}

	printSucceeded()
	err := c.Save()
	if err != nil {
		fmt.Println(err)
	}

	return c
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
