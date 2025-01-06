package search

type ComparableFunc[T any] func(a T, b T) int

/*
Binary search is an efficient algorithm for finding the position of a target element in a sorted array or list.
It works by repeatedly dividing the search interval in half.

The prerequisite for using binary search is that the array or list must be sorted.
If the data is not sorted, you'll need to sort it first, which typically has a time complexity of
O(nlogn). After sorting, you can perform the binary search with a time complexity of O(logn).
*/
func BinarySearch[T any](elements []T, element T, comp ComparableFunc[T]) int {
	start, end := 0, len(elements)
	for start < end {
		mid := start + (end-start)/2
		if comp(elements[mid], element) == 0 {
			return mid
		} else if comp(elements[mid], element) < 0 {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
