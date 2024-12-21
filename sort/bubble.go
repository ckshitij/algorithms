package sort

/*
Bubble Sort is a simple comparison-based sorting algorithm. It repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order. This process continues until no more swaps are needed, indicating that the list is sorted.

How it Works:
- Compare adjacent elements in the list.
- If the first element is greater than the second, swap them.
- Repeat the process for the entire list, moving the largest unsorted element to its correct position with each pass.
- Continue until the entire list is sorted.
*/
func Bubble[T any](nums []T, compare ComparableFunc[T]) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if compare(nums[j], nums[j+1]) {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
