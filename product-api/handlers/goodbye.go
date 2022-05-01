package handlers

import (
	"fmt"
	"log"
	"net/http"
)

//handler
type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// ServeHTTP implements the go http.Handler interface
func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Goodbye request")
	fmt.Fprintf(rw, "Goodbye")
}
