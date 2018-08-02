package mathsy

import "testing"

func TestMultiply(t *testing.T) {
	total := Multiply(5, 5)
	result := 25
	if total != result {
		t.Errorf("Multiply was incorrect, got: %d, want: %d.", total, result)
	}
}
