package digits

import (
	"fmt"
	"slices"
)

// BinDigits erwartet eine Zahl n und liefert eine Liste von Ziffern
func BinDigits(n int) []int {
	result := []int{}

	for n != 0 {
		last_digit := n % 2
		result = append(result, last_digit)
		// result = append([]int{last_digits}, result) //Direkt umgekehrt anhängen.
		n /= 2 // n = n / 2
	}

	slices.Reverse(result)

	return result
}

// HexDigits erwartet eine Zahl n und liefert eine Liste von Ziffern
func HexDigits(n int) []int {
	result := []int{}

	for n != 0 {
		last_digit := n % 16
		result = append(result, last_digit)
		n /= 16 // n = n / 16
	}

	slices.Reverse(result)

	return result
}

// alles in einem: Digits erwartet eine Zahl n und eine Basis
func Digits(n, Base int) []int {
	result := []int{}

	for n != 0 {
		last_digit := n % Base
		result = append(result, last_digit)
		n /= Base
	}

	slices.Reverse(result)

	return result
}

// Sum erwartet eine Liste von Zahlen und berechnet deren Summer.
func Sum(numbers []int) int {
	result := 0
	for _, n := range numbers {
		result += n
	}
	return result
}

// paritätbit erwaret eine Zahl n und liefert :
//
// 1: Falls die Anzahl der Einsen in der Binärdarstellung von n ungerade ist.
// 0: Falls die Anzahl der Einsen in der Binärdarstellung von n gerade ist.
func ParityBit(n int) int {
	return DigitsSum(n, 2) % 2
	//return Sum(Digits(n, 2)) % 2
}

// DigitSum berechnet die Quersumme von n (zur Basis 10).
func DigitsSum(n, base int) int {
	return Sum(Digits(n, base))
}

func ExampleDigits() {
	fmt.Println(Digits(42, 2))
	fmt.Println(Digits(42, 2))
	fmt.Println(Digits(42, 10))
	fmt.Println(Digits(42, 8))

	fmt.Println(ParityBit(42))
	fmt.Println(ParityBit(43))

	//Output:
	//[ 1 0 1 0 1 0 ]
	//[2 10]
	//[4 2]
	//[5 2]
	// 1
	// 0
}
