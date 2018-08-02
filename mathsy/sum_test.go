package mathsy

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	result := 10
	if total != result {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, result)
	}
}

func TestSum2(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}
