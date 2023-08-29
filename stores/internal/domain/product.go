package domain

import "github.com/stackus/errors"

var (
	ErrProductNameIsBlank     = errors.Wrap(errors.ErrBadRequest, "the product name cannot be blank")
	ErrProductPriceIsNegative = errors.Wrap(errors.ErrBadRequest, "the product price cannot be negative")
)

type Product struct {
	ID          string
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

func CreateProduct(id string, storeID string, name string, description string, sku string, price float64) (*Product, error) {
	if name == "" {
		return nil, ErrProductNameIsBlank
	}
	if price < 0 {
        return nil, ErrProductPriceIsNegative
    }
	return &Product{
		ID:          id,
        StoreID:     storeID,
        Name:        name,
        Description: description,
        SKU:         sku,
        Price:       price,
    }, nil
}
