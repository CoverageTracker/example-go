package inventory

// Order is a single-SKU line order awaiting a total price.
type Order struct {
	SKU      string
	Subtotal float64
	Quantity int
	WeightKg float64
	Zone     Zone
	Express  bool
}

// Total applies the volume discount and adds shipping to produce the final
// charge for the order.
func (o Order) Total() (float64, error) {
	if err := ValidateSKU(o.SKU); err != nil {
		return 0, err
	}

	discounted := ApplyDiscount(o.Subtotal, o.Quantity)
	shipping, err := ShippingCost(discounted, o.WeightKg, o.Zone, o.Express)
	if err != nil {
		return 0, err
	}

	return discounted + shipping, nil
}
