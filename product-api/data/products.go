package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"_"`
	UpdatedOn   string  `json:"_"`
	DeletedOn   string  `json:"_"`
}

type Products []*Product

func (p *Product) Validate() error {

	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

func DeleteProduct(id int) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	productList[pos] = &Product{}

	return nil
}

func findProduct(id int) (*Product, int, error) {
	for l, p := range productList {
		if p.ID == id {
			return p, l, nil
		}
	}

	return nil, 0, fmt.Errorf("Product not found")
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
