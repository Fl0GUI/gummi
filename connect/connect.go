package connect

import (
	"log"
	"sync"

	"j322.ica/gumroad-sammi/config"
)

func Connect(config *config.Configuration) {
	wg := sync.WaitGroup{}
	{
		wg.Add(1)
		go func() {
			log.Println("Heartbeat: starting")
			connectHeartbeat(config)
			wg.Done()
		}()
	}
	if config.GumroadConfig.Active {
		wg.Add(1)
		go func() {
			log.Println("Gumroad: starting")
			connectGumroad(config)
			wg.Done()
		}()
	}
	if config.FourthWallConfig.Active {
		wg.Add(1)
		go func() {
			log.Println("Fourthwall: starting")
			connectFourthwall(config)
			wg.Done()
		}()
	}
	if config.ThroneConfig.Active {
		wg.Add(1)
		go func() {
			log.Println("Throne: starting")
			connectThrone(config)
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("Disconnected")
}
