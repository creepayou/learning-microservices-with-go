package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//Hotate is a simple header
type Hotate struct {
	l *log.Logger
}

//NewHotate creates a new Hotate handler with the given logger
func NewHotate(l *log.Logger) *Hotate{
	return &Hotate{l}
}

//ServeHTTP implements the go http.Handler interface
func (h *Hotate) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	h.l.Println("tetaho.")

	//read body
	data, err := io.ReadAll(r.Body)
	if err!= nil{
		http.Error(rw, "Oopsie", http.StatusBadRequest)
		return
	}

	//write response
	fmt.Fprintf(rw, "Echo %s", data)
}