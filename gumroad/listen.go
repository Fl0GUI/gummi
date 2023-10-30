package gumroad

import (
	"embed"
	_ "embed"
	"fmt"
	"log"
	"net/http"
)

var listenUrl = "/sale"

//go:embed test
var testFile embed.FS

type saleHandler struct {
	saleChan chan Sale
}

func (h *saleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("A sale!!")
	err := r.ParseForm()
	if err != nil {
		log.Printf("Got a sale but could not parse the data: %s\n", err)
		return
	}

	h.saleChan <- Sale(r.Form)
}

func (c *Client) Listen(out chan Sale) *http.Server {
	m := http.NewServeMux()
	m.Handle(listenUrl, &saleHandler{out})
	m.Handle("/", http.FileServer(http.FS(testFile)))

	serv := http.Server{}
	serv.Addr = fmt.Sprintf(":%s", c.port)
	serv.Handler = m
	return &serv
}
