package fourthwall

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"j322.ica/gumroad-sammi/config"
)

var listenUrl = "/fourthwall"

var handler *SaleHandler

type SaleHandler struct {
	saleChan chan Sale
	key      []byte
}

func GetSalesChan() chan Sale {
	return handler.saleChan
}

func SetSecretKey(key []byte) {
	handler.key = key
}

func (h *SaleHandler) verify(mac, data *[]byte) error {
	if len(h.key) == 0 {
		return nil
	}
	macHasher := hmac.New(sha256.New, h.key)
	macHasher.Write(*data)
	calcMac := macHasher.Sum(nil)
	if len(calcMac) > len(*mac) {
		return errors.New("Invalid HMac")
	}
	if !hmac.Equal(calcMac, (*mac)[:len(calcMac)]) {
		return errors.New("Invalid HMac")
	}
	return nil
}

func (h *SaleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hmacs := r.Header["X-Fourthwall-Hmac-Sha256"]
	if len(hmacs) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hmac := make([]byte, 2*len(hmacs[0]))
	base64.StdEncoding.Decode(hmac, []byte(hmacs[0]))
	if err := h.verify(&hmac, &data); err != nil {
		log.Println("Fourthwall fake sale: rejected")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var saleData Sale
	err = json.Unmarshal(data, &saleData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

	h.saleChan <- saleData
}

func Handle(c *config.Advanced, m *http.ServeMux) {
	out := make(chan Sale, c.BufferSize)
	handler = &SaleHandler{out, []byte{}}
	m.Handle(listenUrl, handler)
}
