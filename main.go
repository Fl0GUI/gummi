package main

import (
	"fmt"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/connect"
	"j322.ica/gumroad-sammi/gummi"
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
		//gummi.Intro(&funcs, c)
		fmt.Printf("Invalid setup: %s\n", funcs)
		return
	}
	for !funcs.Valid() {
		gummi.Fix(&funcs, c)
		funcs = validate.Validate(c)
	}
	c.Save()

	// connect stores to sammi
	connect.Connect(c)

	// wait for close?
	<-make(chan struct{})
}
