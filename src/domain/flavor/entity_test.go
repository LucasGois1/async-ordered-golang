package flavor

import "testing"

func TestFlavor_String(t *testing.T) {
	tests := []struct {
		name string
		f    Flavor
		want string
	}{
		{name: "Should return the name Vanilla", f: Vanilla, want: "Vanilla"},
		{name: "Should return the name Chocolate", f: Chocolate, want: "Chocolate"},
		{name: "Should return the name Strawberry", f: Strawberry, want: "Strawberry"},
		{name: "Should return the name Mint", f: Mint, want: "Mint"},
		{name: "Should return the name Yogurt", f: Yogurt, want: "Yogurt"},
		{name: "Should return the name Grape", f: Grape, want: "Grape"},
		{name: "Should return the name Lemon", f: Lemon, want: "Lemon"},
		{name: "Should return the name Pistachio", f: Pistachio, want: "Pistachio"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Name(tt.f); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
