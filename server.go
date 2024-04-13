package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"j322.ica/gumroad-sammi/gummi"
	"j322.ica/gumroad-sammi/gumroad"
	"j322.ica/gumroad-sammi/sammi"
)

func runServer() {
	log.Println("gummi: send gumroad sales to SAMMI Core")
	c := gummi.Repl()
	log.Println("Starting gummi application")

	log.Println("Starting SAMMI Core client.")
	sammiC := sammi.NewClient(c)

	log.Println("Starting gumroad client")
	gumroadC := gumroad.NewClient(c)
	salesChan, gumroadErr := gumroadC.Subscribe()
	go func() {
		for err := range gumroadErr {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("Gumroad setup error: %s\n", err)
			}
		}
	}()

	log.Println("Forwarding sales")
	for sale := range salesChan {
		err := pushSale(sammiC, sale)
		if err != nil {
			log.Printf("Error when forwarding sale: %s\n", err)
		} else {
			log.Println("Forwarded a sale!")
		}
	}
}

func pushSale(c *sammi.Client, s gumroad.Sale) error {
	var sammiErrors = make([]error, 0, len(s))
	for k, v := range s {
		key := strings.Replace(strings.Replace(k, "[", "_", -1), "]", "", -1)
		sammiErrors = append(sammiErrors, c.SetVariable(key, v[0]))
	}
	sammiErrors = append(sammiErrors, c.PushButton())
	return errors.Join(sammiErrors...)
}
