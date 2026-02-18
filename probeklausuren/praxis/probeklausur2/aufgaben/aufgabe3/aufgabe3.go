package aufgabe3

import "math"

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * RANDBEDINGUNG: Die Funktion muss rekursiv sein.
 * ERREICHBARE PUNKTE: 10
 */

// CountSquares erwartet eine Liste von Zahlen.
// CountSquares liefert die Anzahl der QuadratzahlenZahlen in der Liste.
func CountSquares(list []int) int {
	if len(list) == 0 {
		return 0
	}
	count := 0

	for i := 0; i < len(list); i++ {
		o := list[0]
		u := math.Sqrt(float64(o))
		if int(u)*int(u) == o {
			count = 1
		}
	}
	return count + CountSquares(list[1:])
}
