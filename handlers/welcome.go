package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Welcome struct {
	l *log.Logger
}

func NewWelcome(l *log.Logger) *Welcome {
	return &Welcome{l}
}

// Welcome handles GET requests at "/"
func (w *Welcome) Welcome(rw http.ResponseWriter, r *http.Request) {
	w.l.Println("Handle Welcome requests")
	fmt.Fprintln(rw, "Welcome to The Office API!")
}
