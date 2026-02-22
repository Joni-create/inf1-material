package mergesort

import "fmt"

// Sortiert die Liste l mit dem Verfahren MergeSort.
func MergeSort(l []int) {
	if len(l) <= 1 {
		return
	}
	mid := len(l) / 2
	left := make([]int, mid)
	right := make([]int, len(l)-mid)
	copy(left, l[:mid])
	copy(right, l[mid:])
	MergeSort(left)
	MergeSort(right)
	merge(left, right, l)
}

func merge(left, right, result []int) {
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
}

func ExampleMergeSort() {
	l1 := []int{17, 25, 22, 3, 15, 4, 35, 105, 42, 1}
	l2 := []int{3, 6, 1, 8, 2, 5, 105, 68, 34, 23, 89, 21, 7, 4}

	MergeSort(l1)
	MergeSort(l2)

	fmt.Println(l1)
	fmt.Println(l2)

	// Output:
	// [1 3 4 15 17 22 25 35 42 105]
	// [1 2 3 4 5 6 7 8 21 23 34 68 89 105]
}
