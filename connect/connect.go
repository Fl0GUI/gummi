package connect

import (
	"fmt"
	"log"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/fourthwall"
	"j322.ica/gumroad-sammi/gumroad"
	"j322.ica/gumroad-sammi/sammi"
)

func Connect(config *config.Configuration) {
	if config.GumroadConfig.Active {
		go connectGumroad(config, config.GumroadConfig.ButtonId)
	}
	if config.FourthWallConfig.Active {
		go connectFourthwall(config, config.FourthWallConfig.ButtonId)
	}
}

func connectGumroad(c *config.Configuration, buttonId string) {
	gc := gumroad.NewClient(c)
	if err := backoff(gc.Subscribe, c); err != nil {
		panic(err)
	}
	bc := sammi.NewButtonClient(&c.SammiConfig, buttonId)
	sales := gumroad.GetChannel()

	for sale := range sales {
		err := sendSale(bc, gumroadToVar(sale), c)
		if err != nil {
			log.Printf("Failed to trigger a gumroad sale: %s\n", err)
		} else {
			log.Println("Got a gumroad sale")
		}
	}
}

func connectFourthwall(c *config.Configuration, buttonId string) {
	sammiC := &c.SammiConfig
	bc := sammi.NewButtonClient(sammiC, buttonId)
	sales := fourthwall.GetSalesChan()

	for sale := range sales {
		err := sendSale(bc, sale, c)
		if err != nil {
			log.Printf("Failed to trigger a fourthwall sale: %s\n", err)
		} else {
			log.Println("Got a fourthwall sale")
		}
	}
}

func sendSale(bc *sammi.ButtonClient, sale map[string]interface{}, c *config.Configuration) error {
	sendVar := func() error {
		return bc.SetVariable("Sale", sale)
	}
	err := backoff(sendVar, c)
	if err != nil {
		return fmt.Errorf("Could not send sale: %w", err)
	}

	activate := func() error {
		return bc.PushButton()
	}
	err = backoff(activate, c)
	if err != nil {
		return fmt.Errorf("Could not send sale: %w", err)
	}
	return err
}
