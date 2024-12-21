# Sorting Algorithms

Sorting algorithms are techniques used to arrange elements of a list or collection in a specific order, typically numerical or lexicographical (alphabetical). These algorithms operate on data structures such as arrays or linked lists and are fundamental in computer science for organizing and optimizing data processing.

## Key Characteristics of Sorting Algorithms

1. **Stability**: A sorting algorithm is stable if it preserves the relative order of elements with equal keys.
2. **Time Complexity**: Indicates the algorithm's efficiency in terms of the number of comparisons or swaps needed as the input size grows (e.g., \(O(n^2)\), \(O(n log n)\)).
3. **Space Complexity**: Reflects the additional memory space required beyond the input data.
4. **In-place vs. Out-of-place**:
   - **In-place** algorithms require a constant amount of extra memory.
   - **Out-of-place** algorithms may require additional data structures for sorting.

## Types of Sorting Algorithms

### Comparison-based Sorting
- Compares elements to determine their order.
- **Examples**:
  - **Bubble Sort**: Simple, stable, \(O(n^2)\) time complexity.
  - **Merge Sort**: Stable, \(O(n log n)\) time complexity.
  - **Quick Sort**: Not stable, \(O(n^2)\) worst-case but \(O(n log n)\) average-case time complexity.
  - **Heap Sort**: Not stable, \(O(n log n)\) time complexity.

### Non-comparison Sorting
- Uses data-specific operations without directly comparing elements.
- **Examples**:
  - **Counting Sort**: Works for integer keys, \(O(n + k)\) time complexity where \(k\) is the range of input.
  - **Radix Sort**: Sorts digits or characters in multiple passes, \(O(d \cdot (n + b))\) time complexity where \(d\) is the number of digits and \(b\) is the base.
  - **Bucket Sort**: Divides elements into buckets, \(O(n)\) average-case for uniformly distributed data.

## Applications

Sorting is integral in:
- **Searching** (e.g., binary search).
- **Data analysis** (e.g., finding medians or duplicates).
- **Algorithms that require pre-sorted input** (e.g., divide-and-conquer techniques).
- **Optimizing other computational processes**.
