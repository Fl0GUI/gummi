package gummi

import (
	"fmt"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/validate"
)

func fixFourthWall(f *validate.Functions, config *config.Configuration) {
	adv := &config.Advanced.ServerConfig
	for {
		fmt.Print(gummiSmile)
		gummiSay("Let's set up and test a fourthwall hook to me.")
		gummiSay("Open up https://my-shop.fourthwall.com/admin/dashboard/settings/for-developers?redirect please.")
		gummiSay("Now press the 'Create webhook' button.")
		gummiSay(fmt.Sprintf("In the input field labeled URL put in 'http://%s:%s/fourthwall'.", adv.PublicIp, adv.ServerPort))
		gummiSay("For the event you can select everything that you want.")
		fmt.Println(gummiSmile)
		gummiSay("On that page you should have seen a secret token")
		gummiSay("Can you paste that here? You can also enter a blank value.")
		config.FourthWallConfig.AccessToken = prompt()
		gummiSay("Now press the 'send test notification' button, and I'll receive it if all is good.")
		gummiSay("I'll wait up to one minute for it")
		err := validate.ValidateFourthWallHook(config)
		if err == nil {
			return
		}
		printFailure()
	}
}
