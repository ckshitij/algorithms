package sort

/*
Selection Sort is a comparison-based sorting algorithm that divides the input list into two parts:
a sorted portion and an unsorted portion.
The algorithm repeatedly selects the smallest (or largest, depending on the order)
element from the unsorted portion and swaps it with the first element in the unsorted portion,
moving it into the sorted portion.

How Selection Sort Works:
Start with the entire list as unsorted.
Find the smallest element in the unsorted portion of the list.
Swap it with the first element of the unsorted portion.
Move the boundary between the sorted and unsorted portions by one element.
Repeat steps 2–4 until the entire list is sorted.
*/
func Selection[T any](elements []T, compare ComparableFunc[T]) {
	n := len(elements)
	for i := 0; i < n; i++ {
		temp := elements[i]
		temp_ind := i
		for j := i; j < n; j++ {
			if compare(elements[j], temp) {
				temp = elements[j]
				temp_ind = j
			}
		}
		elements[i], elements[temp_ind] = elements[temp_ind], elements[i]
	}
}
