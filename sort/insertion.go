package sort

/*
Insertion Sort is a simple and efficient comparison-based sorting algorithm,
particularly effective for small datasets or nearly sorted data.
It works by dividing the list into a sorted portion and an unsorted portion,
then repeatedly taking elements from the unsorted portion and inserting them into their correct position in the sorted portion.

How Insertion Sort Works:
Start with the first element as the "sorted portion."
Pick the next element from the unsorted portion.
Compare it with elements in the sorted portion and shift larger elements one position to the right.
Insert the picked element into its correct position.
Repeat until all elements are sorted.
*/
func Insertion[T any](elements []T, compare ComparableFunc[T]) {
	n := len(elements)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		k := i
		for j := i - 1; j >= 0; j-- {
			if compare(elements[k], elements[j]) {
				elements[k], elements[j] = elements[j], elements[k]
				k = j
			} else {
				break
			}
		}
	}
}
