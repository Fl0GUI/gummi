package gummi

import (
	"fmt"

	"j322.ica/gumroad-sammi/config"
)

func fixGumroad(c *config.Config, f *functions) {
	for !f.gumroadServer {
		fmt.Print(gummiSquint)
		gummiSay("I started a server to listen for your gumroad sales, but I could not connect to it.")
		gummiSay(fmt.Sprintf("It should have been reachable on '%s:%s'.", c.GumroadConfig.PublicIp, c.GumroadConfig.ServerPort))
		gummiSay("This most likely means you have not set up port forwarding on your router.")
		gummiSay("You will have to help me with that. Find the management page of your router and set up a port forward rule.")
		gummiSay(fmt.Sprintf("The rule needs to redirect port %s from outside, to your computer's ip on the same port.", c.GumroadConfig.ServerPort))
		gummiSay(fmt.Sprintf("You should also check the firewall settings on your computer. The public internet needs to access me."))
		gummiSay("You can enter anything when you're done.")
		prompt()
		testSammi(c, f)
	}

	for !f.gumroadToken {
		fmt.Print(gummiSmile)
		gummiSay("I also need an access token to contact gumroad.")
		gummiSay("Gumroad has a tutorial for this: https://help.gumroad.com/article/280-create-application-api#For-just-your-account-Wr9_c")
		gummiSay("Just follow those steps and paste the \"Access Token\" here.")
		c.GumroadConfig.AccessToken = prompt()

		testGumroad(c, f)
	}
}
