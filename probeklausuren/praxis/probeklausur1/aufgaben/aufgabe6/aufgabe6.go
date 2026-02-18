package aufgabe6

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion SymmetricDifference.
MAX. PUNKTE: 10
*/

// SymmetricDifference erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die in einer, aber nicht in beiden Listen vorhanden sind.
//
// Die Elemente kommen dabei in der gleichen Reihenfolge vor, wie in den
// Ursprungslisten, mehrfach vorkommende Elemente werden entsprechend wiederholt.
// Die Elemente aus l1 kommen vor denen aus l2 in der Ergebnisliste vor.
func SymmetricDifference(l1, l2 []int) []int {
	// TODO
	l33 := []int{}
	for i := 0; i < len(l1) && i < len(l2); i++ {
		if !contains(l2, l1[i]) {
			l33 = append(l33, l1[i])
		}
		if !contains(l1, l2[i]) {
			l33 = append(l33, l2[i])
		}
	}
	if len(l1) > len(l2) {
		for i := len(l2); i < len(l1); i++ {
			if !contains(l2, l1[i]) {
				l33 = append(l33, l1[i])
			}
		}
	} else if len(l2) > len(l1) {
		for i := len(l1); i < len(l2); i++ {
			if !contains(l1, l2[i]) {
				l33 = append(l33, l2[i])
			}
		}
	}
	return l33

}

// contains returns true if slice s contains element e.
func contains(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
