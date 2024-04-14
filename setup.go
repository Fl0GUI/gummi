package main

import "j322.ica/gumroad-sammi/gummi"

func setup() {
	gummi.Setup()

	// set up http server
	go listen()
}
