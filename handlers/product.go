package handlers

import (
	"example/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func (p*Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		p.getProducts(rw, r)
		return
	}	

	//catchAll
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p*Products) getProducts(rw http.ResponseWriter, r *http.Request){
	productList := data.GetProducts()
	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}