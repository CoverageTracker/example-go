package inventory

import "testing"

func TestOrder_Total(t *testing.T) {
	o := Order{
		SKU:      "TOY-0042",
		Subtotal: 50,
		Quantity: 5,
		WeightKg: 3,
		Zone:     ZoneContinental,
		Express:  false,
	}
	total, err := o.Total()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantDiscounted := 50.0
	wantShipping := 9.0 + 1.5*3
	want := wantDiscounted + wantShipping
	if total != want {
		t.Errorf("got %v, want %v", total, want)
	}
}

func TestOrder_Total_InvalidSKU(t *testing.T) {
	o := Order{
		SKU:      "bad",
		Subtotal: 200,
		Quantity: 20,
		WeightKg: 3,
		Zone:     ZoneContinental,
	}
	if _, err := o.Total(); err == nil {
		t.Error("expected error for invalid SKU, got nil")
	}
}
