package gummi

import (
	"fmt"

	"j322.ica/gumroad-sammi/config"
)

func printOptions(config *config.Configuration) {
	fmt.Println(gummiYay)
	gummiSay("I will help you connect multiple online stores to your sammi dashboard.")
	gummiSay("I support Gumroad, fourthwall, AND throne.")
	gummiSay("Do you want to set up Gumroad with me?")
	config.GumroadConfig.Active = yesNoPrompt()
	gummiSay("Do you want to set up fourthwall with me?")
	config.FourthWallConfig.Active = yesNoPrompt()
	gummiSay("Do you want to set up throne with me?")
	config.ThroneConfig.Active = yesNoPrompt()
}
