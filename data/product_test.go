package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{
		ID:          0,
		Name:        "Tea",
		Description: "",
		Price:       20,
		SKU:         "abs-agh-zdh",
		CreatedOn:   "",
		UpdatedOn:   "",
		DeletedOn:   "",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
