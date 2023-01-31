package __native_testing

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.00
	expect := 5.0

	result := CalculateTax(amount)

	if result != expect {
		t.Errorf("Expected %f but got %f", expect, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{0, 0},
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)

		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received: %f but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Received: %f but expected 0", result)
		}
	})
}
