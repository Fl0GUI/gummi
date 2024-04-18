package gummi

import (
	"errors"
	"fmt"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/sammi"
	"j322.ica/gumroad-sammi/validate"
)

func fixFourthWall(f *validate.Functions, config *config.Configuration) {
	fixFourthWallSimple(f, config)
	fixFourthWallInteractive(f, config)
}

func fixFourthWallSimple(f *validate.Functions, config *config.Configuration) {
	fourthWallC := &config.FourthWallConfig
	for f.FourthWall != nil {
		if errors.Is(f.FourthWall, sammi.ButtonIdNotFoundError) {
			fixFourthWallButton(f, fourthWallC)
		}
		f.FourthWall = validate.ValidateFourthWall(config)
		if f.FourthWall == nil {
			return
		}
		printFailure()
	}
}

func fixFourthWallInteractive(f *validate.Functions, config *config.Configuration) {
	adv := &config.Advanced.ServerConfig
	for {
		fmt.Print(gummiSmile)
		gummiSay("Let's set up and test a fourthwall hook to me.")
		gummiSay("Open up https://my-shop.fourthwall.com/admin/dashboard/settings/for-developers?redirect please.")
		gummiSay("Now press the 'Create webhook' button.")
		gummiSay(fmt.Sprintf("In the input field labeled URL put in 'http://%s:%s/fourthwall'.\n", adv.PublicIp, adv.ServerPort))
		gummiSay("And you can select everything that you want for event.")
		gummiSay("When you're done you can press the 'send test notification' button, and I'll receive it if all is good.")
		gummiSay("I'll wait up to one minute for it")
		err := validate.ValidateFourthWallHook(config)
		if err == nil {
			return
		}
		printFailure()
	}
}

func fixFourthWallButton(f *validate.Functions, c *config.FourthWallConfig) {
	fmt.Print(gummiSmile)
	gummiSay("I need a buttonID to click when you get a fourthwall sale.")
	gummiSay("Can you set one up, and tell me the buttonID please?")
	c.ButtonId = prompt()
}
