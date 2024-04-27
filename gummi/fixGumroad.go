package gummi

import (
	"errors"
	"fmt"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/gumroad"
	"j322.ica/gumroad-sammi/validate"
)

func fixGumroad(f *validate.Functions, config *config.Configuration) {
	serverC := &config.Advanced.ServerConfig
	gumroadC := &config.GumroadConfig
	f.Gumroad = validate.ValidateGumroad(config)
	for f.Gumroad != nil {
		if errors.Is(f.Gumroad, gumroad.Unauthorized) {
			fixGumroadToken(f, gumroadC)
		} else if errors.Is(f.Gumroad, gumroad.SelfTestFailed{}) {
			fixGumroadServer(f, serverC)
		} else {
			panic(f.Gumroad)
		}
		f.Gumroad = validate.ValidateGumroad(config)
		if f.Gumroad == nil {
			return
		}
		printFailure()
	}
}

func fixGumroadServer(f *validate.Functions, serverC *config.ServerConfig) {
	fmt.Print(gummiSquint)
	gummiSay("I started a server to listen for your gumroad sales, but I could not connect to it.")
	gummiSay(fmt.Sprintf("It should have been reachable on '%s:%s'.", serverC.PublicIp, serverC.ServerPort))
	gummiSay("This most likely means you have not set up port forwarding on your router.")
	gummiSay("You will have to help me with that. Find the management page of your router and set up a port forward rule.")
	gummiSay(fmt.Sprintf("The rule needs to redirect port %s from outside, to your computer's ip on the same port.", serverC.ServerPort))
	gummiSay(fmt.Sprintf("You should also check the firewall settings on your computer. The public internet needs to access me."))
	gummiSay("Let me know when I can try again")
	prompt()
}

func fixGumroadToken(f *validate.Functions, gumroadC *config.GumroadConfig) {
	fmt.Print(gummiSmile)
	gummiSay("I also need an access token to contact gumroad.")
	gummiSay("Gumroad has a tutorial for this: https://help.gumroad.com/article/280-create-application-api#For-just-your-account-Wr9_c")
	gummiSay("Just follow those steps and paste the \"Access Token\" here.")
	gumroadC.AccessToken = prompt()

}
