package tax

import "testing"

// go clean -testcache
// run coverage: go test -coverprofile=coverage.out
// run coverage: go tool cover -html=coverage.out
func TestCalculateTaxBtach(t *testing.T) {
	tests := []struct {
		name     string
		tax      float64
		expected float64
	}{
		{"return tax equal 10 ", 1000, 10},
		{"return tax equal 5 ", 500, 5},
		{"return tax equal 10", 1500, 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CalculateTax(test.tax)
			if result != test.expected {
				t.Errorf("Expected %f, got %f", test.expected, result)
			}
		})
	}
}

// run benchmark: go test -bench=.
// run benchmark: go test -bench=. -run=^#
// run benchmark: go test -bench=. -run=^# -count=10
// run benchmark: go test -bench=. -run=^# -count=10 -benchtime=3s
// run benchmark: go test -bench=. -run=^# -benchmem
// run benchmark: go test -bench=. -run=^# -benchmem -cpuprofile=cpu.out
// run benchmark: go test -bench=. -run=^# -benchmem -cpuprofile=cpu.out -memprofile=mem.out
// run benchmark: go test -bench=. -run=^# -benchmem -cpuprofile=cpu.out -memprofile=mem.out -blockprofile=block.out
// run benchmark: go test -bench=. -run=^# -benchmem -cpuprofile=cpu.out -memprofile=mem.out -blockprofile=block.out -trace=trace.out
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(1000)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(1000)
	}
}

// run fuzzing: go test -fuzz=.
// run fuzzing: go test -fuzz=. -fuzztime=3s
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount <= 0 {
			t.Errorf("Expected %f, got %f", 0, result)
		}
	})
}
