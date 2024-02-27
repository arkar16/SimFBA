package structs

type DepthChartPositionDTO struct {
	Position      string
	Archetype     string
	Score         int
	CollegePlayer CollegePlayer
	NFLPlayer     NFLPlayer
}

type ByDCPosition []DepthChartPositionDTO

func (rp ByDCPosition) Len() int      { return len(rp) }
func (rp ByDCPosition) Swap(i, j int) { rp[i], rp[j] = rp[j], rp[i] }
func (rp ByDCPosition) Less(i, j int) bool {
	return rp[i].Score > rp[j].Score
}
