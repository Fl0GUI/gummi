package gummi

import (
	"errors"
	"fmt"
	"time"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/sammi"
	"j322.ica/gumroad-sammi/validate"
)

func fixSammi(f *validate.Functions, config *config.Configuration) {
	c := &config.SammiConfig
	for {
		if errors.Is(f.Sammi, sammi.ServerError) {
			fixSammiConnection(f, c)
		} else if errors.Is(f.Sammi, sammi.AuthError) {
			fixSammiAuth(f, c)
		} else {
			panic(f.Sammi)
		}

		f.Sammi = validate.ValidateSammi(config)
		if f.Sammi == nil {
			return
		}
		printFailure()
	}
}

func fixSammiConnection(f *validate.Functions, c *config.SammiConfig) {
	for {
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
	}
}

func fixSammiAuth(f *validate.Functions, c *config.SammiConfig) {
	fmt.Print(gummiSad)
	gummiSay("I can connect to SAMMI Core, but I am not allowed in.")
	gummiSay("Can you give me the password to the local server?")
	c.Password = prompt()
}
