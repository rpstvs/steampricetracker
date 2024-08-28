package utils

import (
	"fmt"
	"testing"
)

func TestDailyPriceChange(t *testing.T) {

	cases := []struct {
		input1   float64
		input2   float64
		expected float64
	}{
		{
			input1:   1.0,
			input2:   2.0,
			expected: 50,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			actual := DailyPriceChange(c.input1, c.input2)
			fmt.Printf("t: %v\n", actual)
			if actual != c.expected {
				t.Errorf("parsing failed")
			}

		})

	}
}
