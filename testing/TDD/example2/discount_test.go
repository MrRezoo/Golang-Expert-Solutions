package example2

import "testing"

func TestCalculateDiscount(t *testing.T) {
	tests := []struct {
		userType string
		amount   float64
		want     float64
	}{
		{"regular", 100, 0},
		{"vip", 100, 10},
		{"vip", 500, 50},
	}

	for _, tt := range tests {
		t.Run(tt.userType, func(t *testing.T) {
			got := CalculateDiscount(tt.userType, tt.amount)
			if got != tt.want {
				t.Errorf("CalculateDiscount(%v, %v) = %v; want %v", tt.userType, tt.amount, got, tt.want)
			}
		})
	}
}
