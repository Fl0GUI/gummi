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
			log.Println("Gumroad: starting")
			connectGumroad(config, config.GumroadConfig.ButtonId)
			wg.Done()
		}()
	}
	if config.FourthWallConfig.Active {
		wg.Add(1)
		go func() {
			log.Println("Fourthwall: starting")
			connectFourthwall(config, config.FourthWallConfig.ButtonId)
			wg.Done()
		}()
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
	bc := sammi.NewClient(&c.SammiConfig)
	sales := gumroad.GetChannel()

	for sale := range sales {
		err := backoff(func() error { return bc.Trigger("gumroad", gumroadToVar(sale)) }, c)
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
	bc := sammi.NewClient(sammiC)
	sales := fourthwall.GetSalesChan()
	fourthwall.SetSecretKey([]byte(c.FourthWallConfig.AccessToken))

	for sale := range sales {
		trigger := fourthWallTriggerName(sale)
		err := backoff(func() error {
			return bc.Trigger(trigger, sale)
		}, c)
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
