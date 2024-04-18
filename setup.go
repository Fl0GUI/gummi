package main

import (
	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/gummi"
)

func setup(config *config.Configuration) {
	gummi.Setup()

	// set up http server
	go listen(config)
}
