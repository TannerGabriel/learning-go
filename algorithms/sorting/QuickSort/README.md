# Quicksort

Quicksort is a divide and conquer algorithm. Quicksort first divides a large array into two smaller sub-arrays: the low elements and the high elements. Quicksort can then recursively sort the sub-arrays

![Quicksort](https://upload.wikimedia.org/wikipedia/commons/6/6a/Sorting_quicksort_anim.gif)

The steps are:

1. Pick an element, called a pivot, from the array.
2. Partitioning: reorder the array so that all elements with values less than the pivot come before the pivot, while all elements with values greater than the pivot come after it (equal values can go either way). After this partitioning, the pivot is in its final position. This is called the partition operation.
3. Recursively apply the above steps to the sub-array of elements with smaller values and separately to the sub-array of elements with greater values.

## Complexity

| Name                  | Best            | Average             | Worst               | Memory    | Stable    |
| --------------------- | :-------------: | :-----------------: | :-----------------: | :-------: | :-------: | 
| **Quick sort**        | n&nbsp;log(n)   | n&nbsp;log(n)       | n<sup>2</sup>       | log(n)    | No        | 

## Implementation

- [Quick Sort](https://github.com/TannerGabriel/learning-go/blob/master/algorithms/sorting/QuickSort/quicksort.go)

## References

- [Wikipedia](https://en.wikipedia.org/wiki/Quicksort)
- [GeeksForGeeks](https://www.geeksforgeeks.org/quick-sort/)