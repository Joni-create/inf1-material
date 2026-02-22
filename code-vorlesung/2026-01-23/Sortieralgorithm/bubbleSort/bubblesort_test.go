package bubbleSort

import "fmt"

// Sortiert die Liste l mit dem Verfahren BubbleSort.
func BubbleSort(l []int) {
	n := len(l)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
			}
		}
	}
}

func ExampleBubbleSort() {
	l1 := []int{17, 25, 22, 3, 15, 4, 35, 105, 42, 1}
	l2 := []int{3, 6, 1, 8, 2, 5, 105, 68, 34, 23, 89, 21, 7, 4}

	BubbleSort(l1)
	BubbleSort(l2)

	fmt.Println(l1)
	fmt.Println(l2)

	// Output:
	// [1 3 4 15 17 22 25 35 42 105]
	// [1 2 3 4 5 6 7 8 21 23 34 68 89 105]
}
