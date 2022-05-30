package array

import "testing"

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		result := Sum(numbers)
		expected := 15

		if expected != result {
			t.Errorf("result %d, expected %d, dado %v", result, expected, numbers)
		}
	})

	t.Run("Collection of infinity length", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		result := Sum(numbers)
		expected := 6

		if result != expected {
			t.Errorf("result: %d, expected: %d, data: %v", result, expected, numbers)
		}
	})
}
