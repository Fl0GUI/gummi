package main

import (
	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/gummi"
	"j322.ica/gumroad-sammi/validate"
)

func main() {
	// load config and set defaults
	config.NewConfig()
	if err := config.Config.Save(); err != nil {
		panic(err)
	}

	setup()

	// validate loaded config
	funcs := validate.Validate()
	if !funcs.Valid() {
		gummi.Intro(&funcs)
	}
	for !funcs.Valid() {
		gummi.Fix(&funcs)
		funcs = validate.Validate()
	}
	return

	// use config to set up
	setup()

	// wait for close?
	<-make(chan struct{})
}
