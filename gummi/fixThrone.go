package gummi

import (
	"fmt"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/throne"
	"j322.ica/gumroad-sammi/validate"
)

func fixThrone(f *validate.Functions, c *config.Configuration) {
	for {
		fmt.Print(gummiSmile)
		gummiSay("Time to set up throne!")
		gummiSay("I'm going to hijack the browser source extension.")
		gummiSay("Can you set one up at https://throne.com/profile/integrations/browsersource please?")
		gummiSay("Then copy the url, and give it to me.")
		updateCreatorId(f, c)
		for f.Throne != nil {
			gummiSay("Oops, that seems incorrect. Can you try again?")
			gummiSay("It should be formatted like https://throne.com/stream-alerts/...")
			updateCreatorId(f, c)
		}
		return
	}
}

func updateCreatorId(f *validate.Functions, c *config.Configuration) {
	url := prompt()
	c.ThroneConfig.CreatorId, f.Throne = throne.Extract(url)
	if f.Throne != nil {
		return
	}

	f.Throne = validate.ValidateThrone(c)
}
