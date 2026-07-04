package inventory

import "testing"

func TestVolumeDiscount(t *testing.T) {
	cases := []struct {
		quantity int
		want     float64
	}{
		{5, 0},
		{20, 0.05},
		{50, 0.10},
	}
	for _, c := range cases {
		got := VolumeDiscount(c.quantity)
		if got != c.want {
			t.Errorf("VolumeDiscount(%d) = %v, want %v", c.quantity, got, c.want)
		}
	}
}

func TestApplyDiscount(t *testing.T) {
	got := ApplyDiscount(200, 20)
	want := 200 * 0.95
	if got != want {
		t.Errorf("ApplyDiscount(200, 20) = %v, want %v", got, want)
	}
}
