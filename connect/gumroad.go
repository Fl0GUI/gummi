package connect

import (
	"log"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/gumroad"
	"j322.ica/gumroad-sammi/sammi"
)

func connectGumroad(c *config.Configuration) {
	gc := gumroad.NewClient(c)
	if err := backoff(gc.Subscribe, &c.Advanced.BackoffTimes); err != nil {
		panic(err)
	} else {
		log.Println("Gumroad subscription: success")
	}
	bc := sammi.NewClient(&c.SammiConfig)
	sales := gumroad.GetChannel()

	for sale := range sales {
		err := backoff(func() error { return bc.Trigger("gumroad", gumroadToVar(sale)) }, &c.Advanced.BackoffTimes)
		log.Println("Gumroad sale: received")
		if err != nil {
			log.Printf("Gumroad trigger failure: %s\n", err)
		} else {
			log.Println("Gumroad trigger: success")
		}
	}
	if err := backoff(gc.Unsubscribe, &c.Advanced.BackoffTimes); err != nil {
		panic(err)
	} else {
		log.Println("Gumroad unsubscription: success")
	}
}
