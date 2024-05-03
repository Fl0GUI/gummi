package main

import (
	"os"
	"os/signal"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/connect"
	"j322.ica/gumroad-sammi/fourthwall"
	"j322.ica/gumroad-sammi/gummi"
	"j322.ica/gumroad-sammi/gumroad"
	"j322.ica/gumroad-sammi/throne"
	"j322.ica/gumroad-sammi/validate"
)

func main() {
	// load config and set defaults
	c := config.NewConfig()
	if err := c.Save(); err != nil {
		panic(err)
	}

	setup(c)

	// validate loaded config
	funcs := validate.Validate(c)
	if !funcs.Valid() {
		gummi.Intro(&funcs, c)
	}
	for !funcs.Valid() {
		gummi.Fix(&funcs, c)
		funcs = validate.Validate(c)
	}
	c.Save()

	// wait for close?
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)
	go func() {
		<-signals
		close(gumroad.GetChannel())
		close(fourthwall.GetSalesChan())
		throne.Stop()
	}()

	// connect stores to sammi
	connect.Connect(c)
}
