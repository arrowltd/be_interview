// main.go
package main

import (
	"fmt"
)

// Product represents an item in the inventory.
type Product struct {
	ID       string
	Name     string
	Quantity int
}

// Inventory holds a collection of products.
// HINT: The problem might be in this struct's definition.
type Inventory struct {
	products map[string]Product
}

// NewInventory creates and returns a new inventory.
func NewInventory() Inventory {
	return Inventory{
		products: make(map[string]Product),
	}
}

// AddProduct adds a new product to the inventory.
func (inv *Inventory) AddProduct(p Product) {
	inv.products[p.ID] = p
	fmt.Printf("Added: %s, Initial Stock: %d\n", p.Name, p.Quantity)
}

// IncreaseStock finds a product by its ID and increases its quantity.
// TODO: THIS FUNCTION IS BUGGY. FIX IT.
func (inv *Inventory) IncreaseStock(productID string, amount int) error {
	product, found := inv.products[productID]
	if !found {
		return fmt.Errorf("product with ID %s not found", productID)
	}

	// We increase the quantity on the retrieved product object.
	fmt.Printf("Attempting to increase stock for '%s' from %d by %d...\n", product.Name, product.Quantity, amount)
	product.Quantity += amount

	// Do we need to do anything else here?
	// The developer who wrote this assumed that since 'product' came
	// from the map, the map would now hold the updated version.

	return nil
}

// GetStock returns the current stock for a given product ID.
func (inv *Inventory) GetStock(productID string) (int, error) {
	product, found := inv.products[productID]
	if !found {
		return 0, fmt.Errorf("product with ID %s not found", productID)
	}
	return product.Quantity, nil
}

func main() {
	// 1. Setup
	inventory := NewInventory()
	inventory.AddProduct(Product{ID: "p1", Name: "Laptop", Quantity: 10})

	// 2. The Action - Increase the stock
	err := inventory.IncreaseStock("p1", 5)
	if err != nil {
		fmt.Println("Error updating stock:", err)
		return
	}

	// 3. Verification - Check the final stock
	finalStock, _ := inventory.GetStock("p1")

	fmt.Println("---------------------------------")
	fmt.Printf("Final stock for Laptop: %d\n", finalStock)
	fmt.Println("---------------------------------")
	fmt.Println("Expected Output: Final stock for Laptop: 15")
	fmt.Println("Actual Output:   Final stock for Laptop:", finalStock)

	if finalStock == 15 {
		fmt.Println("\n✅ Test Passed!")
	} else {
		fmt.Println("\n❌ Test Failed! The stock was not updated.")
	}
}
