package sort

/*
Bubble Sort is a simple comparison-based sorting algorithm.
It repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order.
This process continues until no more swaps are needed, indicating that the list is sorted.

How it Works:
- Compare adjacent elements in the list.
- If the first element is greater than the second, swap them.
- Repeat the process for the entire list, moving the largest unsorted element to its correct position with each pass.
- Continue until the entire list is sorted.
*/
func Bubble[T any](elements []T, compare ComparableFunc[T]) {
	n := len(elements)
	for i := 0; i < n-1; i++ {
		swapped := false
		// Perform one pass of bubble sort
		for j := 0; j < n-i-1; j++ {
			if compare(elements[j], elements[j+1]) {
				elements[j], elements[j+1] = elements[j+1], elements[j]
				swapped = true
			}
		}
		// If no elements were swapped, the list is sorted
		if !swapped {
			break
		}
	}
}
