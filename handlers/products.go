package handlers

import (
	"context"
	"log"
	"microservice/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

type KeyProduct struct{}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handel GET Products")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Enable to marshal JSON", http.StatusBadRequest)
	}
}

func (p *Products) PostProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handel POST Products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Enable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handel PUT Products")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err != nil {
		http.Error(rw, "Product not found", http.StatusBadRequest)
	}
}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Enable to marshal JSON", http.StatusBadRequest)
			return
		}

		// Add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
