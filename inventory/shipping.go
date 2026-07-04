// Package inventory implements order pricing for a small mail-order
// warehouse: shipping cost, volume discounts, and SKU validation.
package inventory

import "errors"

// Zone is a shipping destination tier; rates and per-kg surcharges vary by zone.
type Zone string

const (
	ZoneDomestic      Zone = "domestic"
	ZoneContinental   Zone = "continental"
	ZoneInternational Zone = "international"
)

var (
	ErrInvalidWeight = errors.New("weight must be positive")
	ErrUnknownZone   = errors.New("unknown shipping zone")
)

// freeShippingThreshold is the discounted subtotal above which standard
// (non-express) shipping is waived.
const freeShippingThreshold = 75.0

// ShippingCost computes the shipping charge for an order. Weight-based
// surcharges scale with distance, express shipping adds a rush surcharge,
// and orders above freeShippingThreshold ship free unless express.
func ShippingCost(subtotal, weightKg float64, zone Zone, express bool) (float64, error) {
	if weightKg <= 0 {
		return 0, ErrInvalidWeight
	}

	var base float64
	switch zone {
	case ZoneDomestic:
		base = 4.0
	case ZoneContinental:
		base = 9.0
	case ZoneInternational:
		base = 18.0
	default:
		return 0, ErrUnknownZone
	}

	var perKg float64
	switch {
	case weightKg <= 1:
		perKg = 0
	case weightKg <= 5:
		perKg = 1.5
	default:
		perKg = 2.75
	}
	cost := base + perKg*weightKg

	if express {
		cost *= 1.6
	}

	if subtotal >= freeShippingThreshold && !express {
		return 0, nil
	}

	return cost, nil
}
