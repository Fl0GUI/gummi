package main

import (
	"errors"
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
	salesChan := make(chan gumroad.Sale, 10)
	go func() {
		log.Println("Starting gumroad notification server")
		if err := gumroadC.Listen(salesChan).ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			close(salesChan)
			panic(err)
		}
	}()

	log.Println("Forwarding sales")
	for sale := range salesChan {
		pushSale(sammiC, sale)
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
