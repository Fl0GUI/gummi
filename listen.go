package main

import (
	"errors"
	"net/http"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/gumroad"
	"j322.ica/gumroad-sammi/server"
)

func listen() {
	m := http.NewServeMux()
	gumroad.Handle(&config.Config.Advanced, m)

	err := server.ListenAndServe(m, &config.Config.Advanced.ServerConfig)
	if !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
