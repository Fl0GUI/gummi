package connect

import (
	"log"
	"time"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/heartbeat"
	"j322.ica/gumroad-sammi/sammi"
)

var Version string = "dev"
var Revision string = "dev"

func connectHeartbeat(c *config.Configuration) {
	bc := sammi.NewClient(&c.SammiConfig)
	// initial heartbeat
	backoff(func() error { return bc.Trigger("gummi:heartbeat", data()) }, &c.HeartbeatConfig)
	// repeat heartbeat
	beat := time.NewTicker(time.Second * time.Duration(c.HeartbeatConfig.Base))
	defer log.Println("Heartbeat: stopped")
	for {
		select {
		case <-beat.C:
			err := backoff(func() error { return bc.Trigger("gummi:heartbeat", data()) }, &c.Advanced.BackoffTimes)
			if err != nil {
				log.Printf("Heartbeat trigger failure: %s\n", err)
			} else {
				log.Println("Heartbeat")
			}
		case <-heartbeat.Heartbeat:
			return
		}
	}
}

func data() map[string]interface{} {
	now := time.Now()
	return map[string]interface{}{
		"date":         now.Format("2006/01/02 03:04:05"),
		"author":       "Fl_GUI",
		"author_site":  "https://openfl.eu",
		"project":      "gummi",
		"project_site": "https://github.com/Fl0GUI/gummi/",
		"contact":      "https://github.com/Fl0GUI/gummi/issues/new/choosehttps://github.com/Fl0GUI/gummi/issues/new/choose",
		"release":      Version,
		"revision":     Revision,
	}
}
