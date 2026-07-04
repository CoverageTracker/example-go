package inventory

// VolumeDiscount returns the fractional discount earned by ordering the
// given quantity of a single SKU.
func VolumeDiscount(quantity int) float64 {
	switch {
	case quantity >= 100:
		return 0.15
	case quantity >= 50:
		return 0.10
	case quantity >= 20:
		return 0.05
	default:
		return 0
	}
}

// ApplyDiscount returns subtotal after the volume discount for quantity is applied.
func ApplyDiscount(subtotal float64, quantity int) float64 {
	return subtotal * (1 - VolumeDiscount(quantity))
}
