package tax

import "testing"

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
