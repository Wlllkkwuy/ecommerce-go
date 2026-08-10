package products

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidPrice = errors.New("el precio debe ser mayor a cero $")
	ErrInvalidStock = errors.New("el stock no puede ser negativo -")
	ErrEmptyName    = errors.New("el nombre del producto no puede estar vacío")
)

type Product struct {
	id    int
	name  string
	price float64
	stock int
}

func NewProduct(id int, name string, price float64, stock int) (*Product, error) {
	if name == "" {
		return nil, ErrEmptyName
	}
	if price <= 0 {
		return nil, ErrInvalidPrice
	}
	if stock < 0 {
		return nil, ErrInvalidStock
	}

	return &Product{
		id:    id,
		name:  name,
		price: price,
		stock: stock,
	}, nil
}

func (p *Product) ID() int {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) Stock() int {
	return p.stock
}

func (p *Product) SetPrice(newPrice float64) error {
	if newPrice <= 0 {
		return fmt.Errorf("error al actualizar precio de %s: %w", p.name, ErrInvalidPrice)
	}
	p.price = newPrice
	return nil
}

func (p *Product) ReduceStock(quantity int) error {
	if quantity <= 0 {
		return errors.New("la cantidad a reducir debe ser mayor a cero")
	}
	if p.stock < quantity {
		return fmt.Errorf("inventario insuficiente para %s: disponible %d, solicitado %d", p.name, p.stock, quantity)
	}
	p.stock -= quantity
	// 😈
	return nil
}

func ApplyDiscount(p Product, percentage float64) (Product, error) {
	if percentage < 0 || percentage > 100 {
		return Product{}, errors.New("el porcentaje de descuento debe estar entre 0 y 100")
	}
	p.price = p.price * (1 - (percentage / 100))
	return p, nil
}

func FilterProducts(catalog []Product, predicate func(Product) bool) []Product {
	var filtered []Product
	for _, p := range catalog {
		if predicate(p) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}
