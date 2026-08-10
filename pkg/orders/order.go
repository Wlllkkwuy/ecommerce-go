package orders

import (
	"ecommerce-go/pkg/products"
	"errors"
	"fmt"
)

type PricedItem interface {
	Price() float64
	Name() string
}

// 😈

type Order struct {
	id    int
	items []PricedItem
	total float64
}

func NewOrder(id int, items []PricedItem) (*Order, error) {
	if len(items) == 0 {
		return nil, errors.New("no se puede crear una orden sin productos")
	}

	order := &Order{
		id:    id,
		items: items,
	}

	order.calculateTotal()

	//calcuad
	err := order.ProcessStockDeduction()
	if err != nil {
		return nil, err
	}
	// ------------------

	return order, nil
}

func (o *Order) calculateTotal() {
	var total float64
	for _, item := range o.items {
		total += item.Price()
	}
	o.total = total
}

func (o *Order) ID() int {
	return o.id
}

func (o *Order) Total() float64 {
	return o.total
}

func (o *Order) Items() []PricedItem {
	return o.items
}

func (o *Order) ProcessStockDeduction() error {
	for _, item := range o.items {
		if prod, ok := item.(*products.Product); ok {
			err := prod.ReduceStock(1)
			if err != nil {
				return fmt.Errorf("error al genearr la orden %d: %w", o.id, err)
			}
		}
	}
	return nil
}
