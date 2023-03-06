package handlers

import (
	"microservice/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProducts
// Returns a list of products
// responses:
//  201:  noContent

// DeleteProduct delete a product from the databse
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Enable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handel DELTE Product", id)

	err = data.DeleteProduct(id)
	if err != nil {
		http.Error(rw, "Product not found", http.StatusBadRequest)
	}

}
