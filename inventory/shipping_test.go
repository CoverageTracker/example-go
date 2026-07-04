package inventory

import (
	"errors"
	"testing"
)

func TestShippingCost_DomesticLightweight(t *testing.T) {
	cost, err := ShippingCost(20, 0.5, ZoneDomestic, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cost != 4.0 {
		t.Errorf("got %v, want 4.0", cost)
	}
}

func TestShippingCost_ContinentalMidweight(t *testing.T) {
	cost, err := ShippingCost(20, 3, ZoneContinental, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 9.0 + 1.5*3
	if cost != want {
		t.Errorf("got %v, want %v", cost, want)
	}
}

func TestShippingCost_ExpressSurcharge(t *testing.T) {
	cost, err := ShippingCost(20, 0.5, ZoneDomestic, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 4.0 * 1.6
	if cost != want {
		t.Errorf("got %v, want %v", cost, want)
	}
}

func TestShippingCost_FreeAboveThreshold(t *testing.T) {
	cost, err := ShippingCost(100, 2, ZoneDomestic, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cost != 0 {
		t.Errorf("got %v, want 0 (free shipping)", cost)
	}
}

func TestShippingCost_InvalidWeight(t *testing.T) {
	_, err := ShippingCost(20, 0, ZoneDomestic, false)
	if !errors.Is(err, ErrInvalidWeight) {
		t.Errorf("got %v, want ErrInvalidWeight", err)
	}
}
