package connect

import (
	"fmt"
	"log"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/fourthwall"
	"j322.ica/gumroad-sammi/sammi"
)

func connectFourthwall(c *config.Configuration) {
	sammiC := &c.SammiConfig
	bc := sammi.NewClient(sammiC)
	sales := fourthwall.GetSalesChan()
	fourthwall.SetSecretKey([]byte(c.FourthWallConfig.AccessToken))

	for sale := range sales {
		trigger := fourthWallTriggerName(sale)
		err := backoff(func() error {
			return bc.Trigger(trigger, sale)
		}, &c.Advanced.BackoffTimes)
		log.Println("Fourthwall sale: received")
		if err != nil {
			log.Printf("Fourthwall trigger failure: %s\n", err)
		} else {
			log.Println("fourthwall trigger: success")
		}
	}
}

func fourthWallTriggerName(s fourthwall.Sale) string {
	var eventT string
	var t interface{}
	var ok bool
	if t, ok = s["type"]; ok {
		eventT, ok = t.(string)
	}
	if !ok {
		eventT = "null"
	}
	return fmt.Sprintf("fourthwall:%s", eventT)
}
