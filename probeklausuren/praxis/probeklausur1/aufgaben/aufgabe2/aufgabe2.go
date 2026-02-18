package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	firstIndex := -1
	lastIndex := -1

	for i, s := range list {
		if s == first && firstIndex == -1 {
			firstIndex = i
		}
		if s == last && lastIndex == -1 {
			lastIndex = i
		}
	}

	if firstIndex == -1 || lastIndex == -1 || lastIndex < firstIndex {
		return []string{}
	}

	var result []string
	for i, s := range list {
		if i < firstIndex || i > lastIndex {
			result = append(result, s)
		}
	}
	return result
}
