package structs

type UpdateDepthChartDTO struct {
	DepthChartID           int
	UpdatedPlayerPositions []CollegeDepthChartPosition
}

type UpdateNFLDepthChartDTO struct {
	DepthChartID           int
	UpdatedPlayerPositions []NFLDepthChartPosition
}
