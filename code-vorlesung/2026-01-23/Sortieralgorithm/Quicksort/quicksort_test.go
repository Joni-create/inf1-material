package quicksort

import "fmt"

func QuickSort(l []int) {
	// Sonderfall: Die Liste ist leer oder hat nur ein Element.
	if len(l) <= 1 {
		return
	}
	pivot := l[0]
	left := []int{}
	right := []int{}
	// Partitionieren der Liste:
	// Kleinere Elemente als das Pivot nach links, größere nach rechts.
	for _, el := range l[1:] {
		if el < pivot {
			left = append(left, el)
		} else {
			right = append(right, el)
		}
	}
	QuickSort(left)
	QuickSort(right)
	// Elemente in die ursprüngliche Liste zurückkopieren.
	for i, el := range left {
		l[i] = el
	}
	l[len(left)] = pivot
	for i, el := range right {
		l[i+len(left)+1] = el
	}
}

func ExampleQuickSort() {
	l1 := []int{17, 25, 22, 3, 15, 4, 35, 105, 42, 1}
	l2 := []int{3, 6, 1, 8, 2, 5, 105, 68, 34, 23, 89, 21, 7, 4}

	QuickSort(l1)
	QuickSort(l2)

	fmt.Println(l1)
	fmt.Println(l2)

	// Output:
}
