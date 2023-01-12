package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *hello {
	return &hello{l}
}

func (h *hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Print("depency injection Logger\n")
	d, err := io.ReadAll(r.Body)
	// log.Printf("Hi %s", d)
	fmt.Fprintf(w, "Hello %s\n", d)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
}
