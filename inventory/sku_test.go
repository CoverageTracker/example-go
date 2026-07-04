package inventory

import (
	"errors"
	"testing"
)

func TestValidateSKU_Valid(t *testing.T) {
	if err := ValidateSKU("TOY-0042"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestValidateSKU_TooShort(t *testing.T) {
	err := ValidateSKU("TOY-1")
	if !errors.Is(err, ErrSKUTooShort) {
		t.Errorf("got %v, want ErrSKUTooShort", err)
	}
}

func TestValidateSKU_LowercasePrefix(t *testing.T) {
	err := ValidateSKU("toy-0042")
	if !errors.Is(err, ErrSKUBadPrefix) {
		t.Errorf("got %v, want ErrSKUBadPrefix", err)
	}
}

func TestValidateSKU_BadSuffixLength(t *testing.T) {
	err := ValidateSKU("TOY-42")
	if !errors.Is(err, ErrSKUTooShort) {
		t.Errorf("got %v, want ErrSKUTooShort", err)
	}
}
