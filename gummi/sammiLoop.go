package gummi

import (
	"fmt"
	"time"

	"j322.ica/gumroad-sammi/config"
)

func fixSammi(c *config.Config, f *functions) {
	for !f.sammi {
		fmt.Print(gummiSmile)
		gummiSay("I do not have the right ip address and / or port number of your sammi installation.")
		if !debug {
			time.Sleep(time.Millisecond * 200)
		}
		gummiSay("Maybe SAMMI is not running, or maybe the local api is not enabled.")
		gummiSay("Check your SAMMI settings, under the 'SAMMI Local API Settings'.")
		if !debug {
			time.Sleep(time.Second)
		}
		gummiSay("Now, what is the ip address of the computer running SAMMI? If you enter blank I will keep my default.")
		if p := prompt(); p != "" {
			c.Host = p
		}
		gummiSay("Next, what is the port of the local server? Again you can keep this blank for the default.")
		if p := prompt(); p != "" {
			c.Port = p
		}

		testSammi(c, f)
	}

	for !f.sammiAuth {
		fmt.Print(gummiSad)
		gummiSay("I can connect to SAMMI Core, but I am not allowed in.")
		gummiSay("Can you give me the password to the local server?")
		c.Password = prompt()

		testSammi(c, f)
	}

	for !f.sammiButton {
		fmt.Print(gummiSmile)
		gummiSay("I need the buttonID of the button I should trigger. You can find it in the SAMMI Core application. I can not guess a default here.")
		c.ButtonId = prompt()

		testSammi(c, f)
	}
}
