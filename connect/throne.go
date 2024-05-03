package connect

import (
	"log"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/sammi"
	"j322.ica/gumroad-sammi/throne"
)

func connectThrone(c *config.Configuration) {
	if err := throne.Start(&c.ThroneConfig); err != nil {
		panic(err)
	} else {
		log.Println("Throne connection: success")
	}
	bc := sammi.NewClient(&c.SammiConfig)
	sales := throne.GetSalesChan()
	for sale := range sales {
		err := backoff(func() error { return bc.Trigger("throne", sale) }, c)
		if err != nil {
			log.Printf("Throne trigger failure: %s\n", err)
		} else {
			log.Println("Throne trigger: success")
		}
	}
}
