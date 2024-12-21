package sort

import "testing"

func TestSelection(t *testing.T) {
	// Test case for integers
	t.Run("Integers", func(t *testing.T) {
		ints := []int{5, 2, 9, 1, 5, 6}
		expected := []int{1, 2, 5, 5, 6, 9}
		Selection(ints, func(a, b int) bool { return a < b })
		for i, v := range ints {
			if v != expected[i] {
				t.Errorf("Expected %v, got %v", expected, ints)
				break
			}
		}
	})

	// Test case for floats
	t.Run("Floats", func(t *testing.T) {
		floats := []float64{3.1, 2.2, 1.3, 4.4, 0.5}
		expected := []float64{0.5, 1.3, 2.2, 3.1, 4.4}
		Selection(floats, func(a, b float64) bool { return a < b })
		for i, v := range floats {
			if v != expected[i] {
				t.Errorf("Expected %v, got %v", expected, floats)
				break
			}
		}
	})

	// Test case for strings
	t.Run("Strings", func(t *testing.T) {
		strs := []string{"banana", "apple", "cherry", "date"}
		expected := []string{"apple", "banana", "cherry", "date"}
		Selection(strs, func(a, b string) bool { return a < b })
		for i, v := range strs {
			if v != expected[i] {
				t.Errorf("Expected %v, got %v", expected, strs)
				break
			}
		}
	})

	// Test case for custom types
	t.Run("CustomType", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		people := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}
		expected := []Person{
			{"Bob", 25},
			{"Alice", 30},
			{"Charlie", 35},
		}
		Selection(people, func(a, b Person) bool { return a.Age < b.Age })
		for i, v := range people {
			if v != expected[i] {
				t.Errorf("Expected %v, got %v", expected, people)
				break
			}
		}
	})
}
