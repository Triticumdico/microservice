package handlers

import (
	"log"
	"microservice/data"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.postProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("PUT", r.URL.Path)

		reg := regexp.MustCompile("/([0-9+])")
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		ldString := g[0][1]
		id, err := strconv.Atoi(ldString)
		if err != nil {
			p.l.Println("Invalid URI enable to convert to number", ldString)
			http.Error(rw, "Enable to marshal JSON", http.StatusBadRequest)
		}

		p.updateProducts(id, rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handel GET Products")

	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Enable to marshal JSON", http.StatusBadRequest)
	}
}

func (p *Products) postProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handel POST Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Enable to marshal JSON", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handel PUT Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Enable to marshal JSON", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err != nil {
		http.Error(rw, "Product not found", http.StatusBadRequest)
	}
}
