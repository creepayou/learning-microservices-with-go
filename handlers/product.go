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
		p.GetProducts(rw, r)
		return
	}	

	if r.Method == http.MethodPost{
		p.addProduct(rw, r)
		return
	}	

	// if r.Method == http.MethodPut{
	// 	p.updateProduct(rw, r)
	// 	return
	// }	

	//catchAll
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p*Products) GetProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handling GET Request.")
	productList := data.GetProducts()
	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}

func (p*Products) addProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handling POST Request.")

	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil { http.Error(rw, "Unable to decode JSON.", http.StatusBadRequest)}

	data.AddProduct(product)
	productList := data.GetProducts()
	err2 := productList.ToJSON(rw)
	if err2 != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}