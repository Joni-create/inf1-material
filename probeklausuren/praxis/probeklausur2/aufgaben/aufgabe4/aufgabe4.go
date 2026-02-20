package aufgabe4

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// ElementSums erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils die Summe der beiden Elemente enthält.
//
// Annahmen für die Berechnung:
// Falls eine Liste kürzer ist als die andere, soll für die Berechnung der
// hinteren Werte ihr letztes Element verwendet werden.
// Für leere Listen soll für die Berechnung ggf. 0 verwendet werden.
func ElementSums(l1, l2 []int) []int {
	lnew := []int{}
	// TODO
	iMax1 := len(l1)
	l1Max := l1[iMax1-1]
	iMax2 := len(l2)
	l2Max := l2[iMax2-1]
	for i := 0; i < len(l1) || i < len(l2); i++ {

		if i < len(l1) && i < len(l2) {
			lnew = append(lnew, l1[i]+l2[i])
		}

		if i >= len(l2) && i < len(l1) {
			lnew = append(lnew, l1[i]+l2Max)
		}
		if i >= len(l1) && i < len(l2) {
			lnew = append(lnew, l2[i]+l1Max)
		}
	}
	return lnew
}
