package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func intComparator(a, b int) int {
	return a - b
}

func stringComparator(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// User-defined type
type User struct {
	ID   int
	Name string
}

func userComparator(a, b User) int {
	return a.ID - b.ID
}

func TestBinarySearch_Ints(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		element  int
		expected int
	}{
		{
			name:     "Element exists in the middle",
			elements: []int{1, 2, 3, 4, 5},
			element:  3,
			expected: 2,
		},
		{
			name:     "Element exists at the start",
			elements: []int{1, 2, 3, 4, 5},
			element:  1,
			expected: 0,
		},
		{
			name:     "Element exists at the end",
			elements: []int{1, 2, 3, 4, 5},
			element:  5,
			expected: 4,
		},
		{
			name:     "Element does not exist",
			elements: []int{1, 2, 3, 4, 5},
			element:  6,
			expected: -1,
		},
		{
			name:     "Empty list",
			elements: []int{},
			element:  1,
			expected: -1,
		},
		{
			name:     "Single element - match",
			elements: []int{10},
			element:  10,
			expected: 0,
		},
		{
			name:     "Single element - no match",
			elements: []int{10},
			element:  20,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.elements, tt.element, intComparator)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBinarySearch_Strings(t *testing.T) {
	tests := []struct {
		name     string
		elements []string
		element  string
		expected int
	}{
		{
			name:     "Element exists",
			elements: []string{"apple", "banana", "cherry", "date"},
			element:  "cherry",
			expected: 2,
		},
		{
			name:     "Element does not exist",
			elements: []string{"apple", "banana", "cherry", "date"},
			element:  "fig",
			expected: -1,
		},
		{
			name:     "Empty list",
			elements: []string{},
			element:  "apple",
			expected: -1,
		},
		{
			name:     "Single element - match",
			elements: []string{"orange"},
			element:  "orange",
			expected: 0,
		},
		{
			name:     "Single element - no match",
			elements: []string{"orange"},
			element:  "apple",
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.elements, tt.element, stringComparator)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBinarySearch_Users(t *testing.T) {
	tests := []struct {
		name     string
		elements []User
		element  User
		expected int
	}{
		{
			name:     "User exists",
			elements: []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			element:  User{ID: 2},
			expected: 1,
		},
		{
			name:     "User does not exist",
			elements: []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			element:  User{ID: 4},
			expected: -1,
		},
		{
			name:     "Empty list",
			elements: []User{},
			element:  User{ID: 1},
			expected: -1,
		},
		{
			name:     "Single user - match",
			elements: []User{{1, "Alice"}},
			element:  User{ID: 1},
			expected: 0,
		},
		{
			name:     "Single user - no match",
			elements: []User{{1, "Alice"}},
			element:  User{ID: 2},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.elements, tt.element, userComparator)
			assert.Equal(t, tt.expected, result)
		})
	}
}
