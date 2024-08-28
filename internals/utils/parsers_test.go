package utils

import (
	"fmt"
	"testing"
)

func TestSticker(t *testing.T) {

	cases := []struct {
		expected Sticker
		input    []string
		input2   float64
	}{
		{
			expected: Sticker{
				Name:       "Gold Web",
				Tournament: "none",
				Condition:  "none",
				Price:      1.0,
			},
			input:  []string{"Sticker", "Gold Web"},
			input2: 1.0,
		},
		{
			expected: Sticker{
				Name:       "Gold Web",
				Tournament: "none",
				Condition:  "Foil",
				Price:      1.0,
			},
			input:  []string{"Sticker", "Gold Web (Foil) "},
			input2: 1.0,
		},
		{
			expected: Sticker{
				Name:       "Gold Web",
				Tournament: "none",
				Condition:  "Foil",
				Price:      1.0,
			},
			input:  []string{"Sticker", "Gold Web (Foil)"},
			input2: 1.0,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			actual := ParseSticker(c.input, c.input2)
			//fmt.Println(actual)
			if actual != c.expected {
				t.Errorf("parsing failed")
			}

		})

	}
}

func TestSkin(t *testing.T) {

	cases := []struct {
		expected Skin
		input    []string
		input2   float64
	}{
		{
			expected: Skin{
				GunName:   "CZ-75",
				SkinName:  "Gold Web",
				Condition: "field-tested",
				Price:     10.0,
			},
			input:  []string{"CZ-75", "Gold Web (field-tested)"},
			input2: 10.0,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			actual := parseSkin(c.input, c.input2)
			//fmt.Println(actual)
			if actual != c.expected {
				t.Errorf("parsing failed")
			}

		})

	}
}

func TestPrice(t *testing.T) {

	cases := []struct {
		expected float64
		input    string
	}{
		{
			expected: 0.03,
			input:    "0,03",
		},
		{
			expected: 1,
			input:    "1,--",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			actual := PriceConverter(c.input)
			//fmt.Printf("t: %v\n", actual)
			if actual != c.expected {
				t.Errorf("parsing failed")
			}

		})

	}
}
