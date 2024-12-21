# Algorithms Library

A Go library that provides a collection of commonly used algorithms, designed to simplify and optimize solutions to everyday problems in the industry. This library is modular, generic, and easy to use, offering robust implementations for sorting, searching, and other essential computational tasks.

---

## Features

- **Generic Support**: Implements algorithms using Go generics for flexibility and type safety.
- **Industry-Relevant**: Focuses on algorithms widely used in real-world applications.
- **Modular Design**: Each algorithm is organized in its own package for better maintainability.

---

## Algorithms Available

### 1. Sorting
Efficient and reusable sorting algorithms for different use cases:
- **Bubble Sort**: A simple and stable sorting algorithm.
- **Merge Sort**: A divide-and-conquer algorithm with \(O(n log n)\) complexity.
- **Quick Sort**: An efficient in-place sorting technique.

### 2. Searching
Fast and reliable searching algorithms:
- **Binary Search**: Efficiently locate elements in sorted collections.

### 3. Other Algorithms

---

## Installation

To add this module to your project, run:

```bash
go get github.com/ckshitij/algorithms
```

## Example 

```go

package main

import (
	"fmt"
	"github.com/ckshitij/algorithms/sort"
)

func main() {
	// Example: Sorting integers using Bubble Sort
	nums := []int{5, 3, 8, 4, 2}
	sort.Bubble(nums, func(a, b int) bool {
		return a > b // Sort in ascending order
	})
	fmt.Println("Sorted numbers:", nums)

    type Person struct {
        Name string
        Age  int
    }

    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }

    sort.Bubble(people, func(a, b Person) bool {
        return a.Age > b.Age
    })

    fmt.Println("Sorted peoples:", people)
}
```

## Contributing
We welcome contributions! If you have suggestions, bug reports, or feature requests, feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
