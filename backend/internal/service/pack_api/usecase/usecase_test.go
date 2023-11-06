package usecase

import (
	"reflect"
	"testing"
)

func TestCalculatePacksNumber(t *testing.T) {
	tests := []struct {
		itemsOrdered int
		expected     map[int]int
	}{
		{itemsOrdered: 1, expected: map[int]int{250: 1}},
		{itemsOrdered: 250, expected: map[int]int{250: 1}},
		{itemsOrdered: 251, expected: map[int]int{500: 1}},
		{itemsOrdered: 499, expected: map[int]int{500: 1}},
		{itemsOrdered: 501, expected: map[int]int{500: 1, 250: 1}},
		{itemsOrdered: 751, expected: map[int]int{1000: 1}},
		{itemsOrdered: 999, expected: map[int]int{1000: 1}},
		{itemsOrdered: 8750, expected: map[int]int{5000: 1, 2000: 1, 1000: 1, 500: 1, 250: 1}},
		{itemsOrdered: 8751, expected: map[int]int{5000: 2}},
		{itemsOrdered: 12001, expected: map[int]int{5000: 2, 2000: 1, 250: 1}},
	}

	uc := New([]int{250, 5000, 2000, 1000, 500})
	for _, test := range tests {
		actual := uc.CalculatePacksNumber(test.itemsOrdered)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Maps are not equal. Expected: %v, got: %v", test.expected, actual)
		}
	}
}
