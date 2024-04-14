package server

import (
	"embed"
	"fmt"
	"net/http"

	"j322.ica/gumroad-sammi/config"
)

//go:embed test
var testFile embed.FS

func ListenAndServe(mux *http.ServeMux, c *config.ServerConfig) error {
	mux.Handle("/", http.FileServer(http.FS(testFile)))
	serv := http.Server{}
	serv.Addr = fmt.Sprintf(":%s", c.ServerPort)
	serv.Handler = mux
	return serv.ListenAndServe()
}
