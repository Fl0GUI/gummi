package main

import (
	"errors"
	"net/http"

	"j322.ica/gumroad-sammi/config"
	"j322.ica/gumroad-sammi/fourthwall"
	"j322.ica/gumroad-sammi/gumroad"
	"j322.ica/gumroad-sammi/server"
)

func listen(config *config.Configuration) {
	m := http.NewServeMux()
	gumroad.Handle(&config.Advanced, m)
	fourthwall.Handle(&config.Advanced, m)

	err := server.ListenAndServe(m, &config.Advanced.ServerConfig)
	if !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
