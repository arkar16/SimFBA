package structs

type WeatherRegion struct {
	RegionName      string
	ReferenceWeek   string
	Week            uint
	MinLow          int
	AvgLow          int
	MaxLow          int
	MinHigh         int
	AvgHigh         int
	MaxHigh         int
	Clear           float64
	MostlyClear     float64
	PartlyCloudy    float64
	MostlyCloudy    float64
	Overcast        float64
	DaysOfRain      float64
	DaysOfMix       float64
	DaysOfSnow      float64
	MonthlyRainfall float64
	MonthlySnowfall float64
	InchesPerRain   float64
	InchesPerSnow   float64
	WindSpeedMin    float64
	WindSpeedAvg    float64
	WindSpeedMax    float64
}
