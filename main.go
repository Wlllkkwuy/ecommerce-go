package main

import (
	"ecommerce-go/pkg/orders"
	"ecommerce-go/pkg/products"
	"fmt"
)

func main() {
	// Colección inicial de datos
	catalog := []products.Product{
		{ID: 1, Name: "Laptop Gaming", Price: 1200.00, Stock: 5},
		{ID: 2, Name: "Mouse Inalámbrico", Price: 25.00, Stock: 20},
		{ID: 3, Name: "Teclado Mecánico", Price: 80.00, Stock: 10},
	}

	// Uso de Función de Orden Superior pasándole una Función Anónima
	expensiveProducts := products.FilterProducts(catalog, func(p products.Product) bool {
		return p.Price > 50
	})

	// Instanciación de la Orden
	newOrder := orders.CreateOrder(101, expensiveProducts)

	// Salida por consola utilizando el paquete nativo 'fmt'
	fmt.Println("======================================================")
	fmt.Println("                                                 ")
	fmt.Println("   🛒   SISTEMA DE GESTIÓN DE E-COMMERCE (GO)   🛒  ")
	fmt.Println("                                                 ")
	fmt.Println("======================================================")
	fmt.Println("                                                     ")
	fmt.Printf("Orden ID: %d procesada exitosamente.✅\n", newOrder.ID)
	fmt.Printf("Cantidad de items filtrados: %d\n", len(newOrder.Products))
	fmt.Printf("Total acumulado de la orden: $%.2f\n", newOrder.Total)
	fmt.Println("                                                     ")
	fmt.Println("==========================================")
}
