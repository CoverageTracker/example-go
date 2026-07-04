// Command example-go prints the priced total for a small sample order,
// exercising the inventory package end to end.
package main

import (
	"fmt"
	"log"

	"github.com/CoverageTracker/example-go/inventory"
)

func main() {
	order := inventory.Order{
		SKU:      "TOY-0042",
		Subtotal: 249.99,
		Quantity: 30,
		WeightKg: 2.4,
		Zone:     inventory.ZoneContinental,
		Express:  false,
	}

	total, err := order.Total()
	if err != nil {
		log.Fatalf("pricing order: %v", err)
	}

	fmt.Printf("order %s total: $%.2f\n", order.SKU, total)
}
