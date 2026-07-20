package orders

import "ecommerce-go/pkg/products" // Aquí es donde se importa del paquete local

// Struct: Representación del modelo de la Orden de compra
type Order struct {
	ID       int
	Products []products.Product
	Total    float64
}

// Función Pura: Calcula el total acumulativo a partir del slice de productos
func CalculateTotal(items []products.Product) float64 {
	total := 0.0
	for _, item := range items {
		total += item.Price
	}
	return total
}

// Constructor Funcional: Instancia la Orden con el total calculado
func CreateOrder(id int, items []products.Product) Order {
	return Order{
		ID:       id,
		Products: items,
		Total:    CalculateTotal(items),
	}
}
