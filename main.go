package main

// Carreño Wlllkkwuy😈

import (
	"fmt"
	"log"

	"ecommerce-go/pkg/orders"
	"ecommerce-go/pkg/products"
)

func main() {
	fmt.Println("=========================================================")
	fmt.Println("   🛒  SISTEMA DE GESTIÓN DE E-COMMERCE (POO & GO) 🛒    ")
	fmt.Println("=========================================================")

	// 1		Inventario con los precios y productos actualizados

	p1, err := products.NewProduct(1, "💻  Laptop Lenovo V15", 650.00, 6)
	if err != nil {
		log.Fatalf("Error p1: %v", err)
	}

	p2, err := products.NewProduct(2, "📺  Smart TV TCL 50\" QLED 4K", 400.00, 4)
	if err != nil {
		log.Fatalf("Error p2: %v", err)
	}

	p3, err := products.NewProduct(3, "📱  Xiaomi Poco X6 Pro 5G", 380.00, 10)
	if err != nil {
		log.Fatalf("Error p3: %v", err)
	}

	p4, err := products.NewProduct(4, "📱  Samsung Galaxy A35 5G", 320.00, 8)
	if err != nil {
		log.Fatalf("Error p4: %v", err)
	}

	p5, err := products.NewProduct(5, "🔋  EcoFlow River 2 (256Wh)", 270.00, 5)
	if err != nil {
		log.Fatalf("Error p5: %v", err)
	}

	p6, err := products.NewProduct(6, "📱  Redmi Note 13 Pro 5G", 260.00, 12)
	if err != nil {
		log.Fatalf("Error p6: %v", err)
	}

	p7, err := products.NewProduct(7, "🖨️  Impresora Epson L3250", 255.00, 7)
	if err != nil {
		log.Fatalf("Error p7: %v", err)
	}

	p8, err := products.NewProduct(8, "⌚  Amazfit Bip 5", 90.00, 15)
	if err != nil {
		log.Fatalf("Error p8: %v", err)
	}

	p9, err := products.NewProduct(9, "🎧  SoundPEATS Air4", 50.00, 20)
	if err != nil {
		log.Fatalf("Error p9: %v", err)
	}

	// 		Carga del inventario completo

	catalog := []products.Product{*p1, *p2, *p3, *p4, *p5, *p6, *p7, *p8, *p9}

	fmt.Println("\n--- CATÁLOGO DE INVENTARIO COMPLETO ---")
	for _, item := range catalog {
		fmt.Printf("ID %d | %-26s | Precio: $%7.2f | Stock: %2d uds.\n",
			item.ID(), item.Name(), item.Price(), item.Stock())
	}

	// 2 		Programación Funcional = Filtrar productos con precio mayor a $300

	fmt.Println("\n--- FILTRADO FUNCIONAL (Productos Mayores a $300) ---")
	highValueProducts := products.FilterProducts(catalog, func(p products.Product) bool {
		return p.Price() > 300.0
	})

	for _, item := range highValueProducts {
		fmt.Printf("• %-26s -> $%7.2f\n", item.Name(), item.Price())
	}

	// 3 		Uso de Interfaces y Polimorfismo = Simulación de compra tipo "Combo dispositivos"

	fmt.Println("\n--- PROCESAMIENTO DE ORDEN DE COMPRA (Combo amarillo) ---")

	// 		Selección de artículos para el carrito usando la interfaz PricedItem

	var orderItems []orders.PricedItem = []orders.PricedItem{
		p3, // Poco X6 Pro de $380
		p8, // Amazfit bip de 5 $90
		p9, // Soundpeats air 4 de $50
	}

	newOrder, err := orders.NewOrder(201, orderItems)
	if err != nil {
		fmt.Printf("​❗​  Error al generar la orden: %v\n", err)
		return
	}

	fmt.Printf("Orden #%d procesada con éxito.  ✅\n", newOrder.ID())
	fmt.Println("Artículos incluidos en el carrito:")
	for idx, item := range newOrder.Items() {
		fmt.Printf("  %d. %-26s | $%6.2f\n", idx+1, item.Name(), item.Price())
	}
	fmt.Printf("Monto Total Acumulado: $%.2f\n", newOrder.Total())

	// 4 		Muestra  de encapsulación, Reducción de inventario y control de Errores o POO

	fmt.Println("\n--- PRUEBAS DE SEGURIDAD Y VALIDACIÓN (POO) ---")

	// Intento de dar precio malo

	err = p3.SetPrice(-150)
	if err != nil {
		fmt.Printf("[Capturado por Setter]: %v\n", err)
	}

	// 		Disminución de inventario al confirmar la compra

	err = p3.ReduceStock(1)
	if err == nil {
		fmt.Printf("[Stock Actualizado]: A %s le quedan %d unidades.\n", p3.Name(), p3.Stock())
	}

	// 		Intento de sobredemanda de inventario en la tv

	err = p2.ReduceStock(10)
	if err != nil {
		fmt.Printf("[Validación de Stock Excedido]: %v\n", err)
	}

	fmt.Println("\n=======================================================")
	fmt.Println("         ✅   PROCESO FINALIZADO CON ÉXITO   ✅       ")
	fmt.Println("=========================================================")
}
