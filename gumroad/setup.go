package gumroad

import (
	_ "embed"
	"log"
	"net/http"

	"j322.ica/gumroad-sammi/config"
)

var listenUrl = "/gumroad/{pathsecret}"

var handler SaleHandler

type SaleHandler struct {
	saleChan chan Sale
}

func (h *SaleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := r.ParseForm()
	if err != nil {
		log.Printf("Gumroad sale failure: parse failure: %s\n", err)
		return
	}

	h.saleChan <- Sale(r.Form)
}

func Handle(c *config.Advanced, m *http.ServeMux) {
	out := make(chan Sale, c.BufferSize)
	handler = SaleHandler{out}
	m.Handle(listenUrl, &handler)
}

func GetChannel() chan Sale {
	return handler.saleChan
}
