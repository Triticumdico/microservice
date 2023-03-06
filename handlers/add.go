package handlers

import (
	"microservice/data"
	"net/http"
)

func (p *Products) PostProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handel POST Products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
