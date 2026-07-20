package products

// Struct: Definición del tipo de dato Product
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

// Función Pura: Retorna una nueva copia con descuento sin que modifique el producto original
func ApplyDiscount(p Product, percentage float64) Product {
	p.Price = p.Price - (p.Price * (percentage / 100))
	return p
}

// Función de Orden Superior: Recibe un predicado (función anonima) para filtrar el slice
func FilterProducts(products []Product, predicate func(Product) bool) []Product {
	var filtered []Product
	for _, p := range products {
		if predicate(p) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}
