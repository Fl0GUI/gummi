package fourthwall

import (
	"encoding/json"
	"net/http"

	"j322.ica/gumroad-sammi/config"
)

var listenUrl = "/fourthwall"

var handler *SaleHandler

type SaleHandler struct {
	saleChan chan Sale
}

func GetSalesChan() chan Sale {
	return handler.saleChan
}

func (h *SaleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data Sale
	err := json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

	h.saleChan <- data
}

func Handle(c *config.Advanced, m *http.ServeMux) {
	out := make(chan Sale, c.BufferSize)
	handler = &SaleHandler{out}
	m.Handle(listenUrl, handler)
}
