package connect

import (
	"fmt"
	"log"
	"sync"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/fourthwall"
	"j322.ica/gumroad-sammi/gumroad"
	"j322.ica/gumroad-sammi/sammi"
)

func Connect(config *config.Configuration) {
	wg := sync.WaitGroup{}
	if config.GumroadConfig.Active {
		wg.Add(1)
		go func() {
			connectGumroad(config, config.GumroadConfig.ButtonId)
			wg.Done()
		}()
	}
	if config.FourthWallConfig.Active {
		wg.Add(1)
		go connectFourthwall(config, config.FourthWallConfig.ButtonId)
		wg.Done()
	}
	wg.Wait()
}

func connectGumroad(c *config.Configuration, buttonId string) {
	gc := gumroad.NewClient(c)
	if err := backoff(gc.Subscribe, c); err != nil {
		panic(err)
	} else {
		log.Println("Gumroad subscription: success")
	}
	bc := sammi.NewButtonClient(&c.SammiConfig, buttonId)
	sales := gumroad.GetChannel()

	for sale := range sales {
		err := sendSale(bc, gumroadToVar(sale), c)
		log.Println("Gumroad sale: received")
		if err != nil {
			log.Printf("Gumroad trigger failure: %s\n", err)
		} else {
			log.Println("Gumroad trigger: success")
		}
	}
	if err := backoff(gc.Unsubscribe, c); err != nil {
		panic(err)
	} else {
		log.Println("Gumroad unsubscription: success")
	}
}

func connectFourthwall(c *config.Configuration, buttonId string) {
	sammiC := &c.SammiConfig
	bc := sammi.NewButtonClient(sammiC, buttonId)
	sales := fourthwall.GetSalesChan()

	for sale := range sales {
		err := sendSale(bc, sale, c)
		log.Println("Fourthwall sale: received")
		if err != nil {
			log.Printf("Fourthwall trigger failure: %s\n", err)
		} else {
			log.Println("fourthwall trigger: success")
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
