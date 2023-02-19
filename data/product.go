package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"_"`
	UpdatedOn   string  `json:"_"`
	DeletedOn   string  `json:"_"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milk coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
		DeletedOn:   "",
	},
	&Product{
		ID:          2,
		Name:        "Expresso",
		Description: "short an strong coffee",
		Price:       1.29,
		SKU:         "abc124",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
		DeletedOn:   "",
	},
}
