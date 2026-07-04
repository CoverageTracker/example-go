package inventory

import "errors"

var (
	ErrSKUTooShort  = errors.New("sku too short")
	ErrSKUBadPrefix = errors.New("sku must start with a 3-letter category prefix followed by '-'")
	ErrSKUBadSuffix = errors.New("sku must end with a 4-digit numeric code")
)

// ValidateSKU checks that sku matches the warehouse format: a 3-letter
// uppercase category prefix, a hyphen, and a 4-digit numeric code, e.g. "TOY-0042".
func ValidateSKU(sku string) error {
	if len(sku) < 8 {
		return ErrSKUTooShort
	}

	prefix := sku[:3]
	for _, r := range prefix {
		if r < 'A' || r > 'Z' {
			return ErrSKUBadPrefix
		}
	}
	if sku[3] != '-' {
		return ErrSKUBadPrefix
	}

	suffix := sku[4:]
	if len(suffix) != 4 {
		return ErrSKUBadSuffix
	}
	for _, r := range suffix {
		if r < '0' || r > '9' {
			return ErrSKUBadSuffix
		}
	}

	return nil
}
