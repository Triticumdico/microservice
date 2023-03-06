package handlers

import (
	"microservice/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//  200:  productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handel GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serilize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Enable to marshal JSON", http.StatusInternalServerError)
	}
}
