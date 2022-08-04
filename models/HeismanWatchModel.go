package models

type HeismanWatchModel struct {
	FirstName string
	LastName  string
	Position  string
	Archetype string
	School    string
	Score     float64
	Games     int
}

// Sorting Funcs
type ByScore []HeismanWatchModel

func (h ByScore) Len() int      { return len(h) }
func (h ByScore) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h ByScore) Less(i, j int) bool {
	return h[i].Score > h[j].Score
}
