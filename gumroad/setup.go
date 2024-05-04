package gumroad

import (
	_ "embed"
	"io"
	"log"
	"net/http"

	"j322.ica/gumroad-sammi/config"
)

var listenUrl = "/gumroad/{pathsecret}"

var handler SaleHandler

type SaleHandler struct {
	saleChan chan Sale
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "404 page not found\n")
}

func validateSecret(r *http.Request) error {
	secret := r.PathValue("pathsecret")
	return client.isMyUrlSecret(secret)
}

func (h *SaleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := validateSecret(r); err != nil {
		notFoundHandler(w, r)
		log.Printf("Gumroad sale blocked: %s\n", err)
		return
	}
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
	m.HandleFunc("/gumroad", notFoundHandler)
}

func GetChannel() chan Sale {
	return handler.saleChan
}
