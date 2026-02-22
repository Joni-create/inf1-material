package insertionsort

import "fmt"

// Sortiert die Liste l mit dem Verfahren InsertionSort.
func InsertionSort(l []int) {
	n := len(l)
	for i := 1; i < n; i++ {
		key := l[i]
		j := i - 1
		for j >= 0 && l[j] > key {
			l[j+1] = l[j]
			j--
		}
		l[j+1] = key
	}
}

func ExampleInsertionSort() {
	l1 := []int{17, 25, 22, 3, 15, 4, 35, 105, 42, 1}
	l2 := []int{3, 6, 1, 8, 2, 5, 105, 68, 34, 23, 89, 21, 7, 4}

	InsertionSort(l1)
	InsertionSort(l2)

	fmt.Println(l1)
	fmt.Println(l2)

	// Output:
	// [1 3 4 15 17 22 25 35 42 105]
	// [1 2 3 4 5 6 7 8 21 23 34 68 89 105]
}
