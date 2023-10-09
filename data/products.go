package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ProductId int	 	`json:"productId"`
	Name      string	`json:"name"`
	Desc      string	`json:"desc"`
	Price     float32	`json:"price"`
	CreatedOn string	`json:"createdOn"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error{
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetProducts() Products{
	return productList
}


var productList = []*Product{
	{
		ProductId: 1,
		Name:      "Japanese Porridge",
		Desc: `Okayu is a comforting and warming rice 
					porridge dish made with minimal ingredients, 
					it's perfect for those days when you're feeling 
					under the weather and looking for a simple dish 
					to pick you back up!`,
		Price:     3.21,
		CreatedOn: time.Now().UTC().String(),
	},
	{
		ProductId: 2,
		Name:      "X-Potato",
		Desc:      "Potato, but X.",
		Price:     10.10,
		CreatedOn: time.Now().UTC().String(),
	},
}