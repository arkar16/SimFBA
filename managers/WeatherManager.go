package managers

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"

	"github.com/CalebRose/SimFBA/structs"
	"github.com/CalebRose/SimFBA/util"
)

func GenerateWeatherForGames() {

}

func getRegionalWeather() []structs.WeatherRegion {
	path := "C:\\Users\\ctros\\go\\src\\github.com\\CalebRose\\SimFBA\\data\\WeatherData"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var regions []structs.WeatherRegion

	for _, file := range files {
		filePath := path + "\\" + file.Name() + ".csv"
		f, err := os.Open(filePath)
		if err != nil {
			log.Fatal("Unable to read input file "+filePath, err)
		}
		defer f.Close()

		csvReader := csv.NewReader(f)
		rows, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal("Unable to parse file as CSV for "+filePath, err)
		}

		var regionName string = ""

		for idx, row := range rows {

			if idx < 4 {
				if idx == 0 {
					regionName = row[1]
				}
				continue
			}
			region := structs.WeatherRegion{
				RegionName:      regionName,
				ReferenceWeek:   row[0],
				Week:            uint(util.ConvertStringToInt(row[1])),
				MinLow:          util.ConvertStringToInt(row[2]),
				AvgLow:          util.ConvertStringToInt(row[3]),
				MaxLow:          util.ConvertStringToInt(row[4]),
				MinHigh:         util.ConvertStringToInt(row[5]),
				AvgHigh:         util.ConvertStringToInt(row[6]),
				MaxHigh:         util.ConvertStringToInt(row[7]),
				Clear:           util.ConvertStringToFloat(row[8]),
				MostlyClear:     util.ConvertStringToFloat(row[9]),
				PartlyCloudy:    util.ConvertStringToFloat(row[10]),
				MostlyCloudy:    util.ConvertStringToFloat(row[11]),
				Overcast:        util.ConvertStringToFloat(row[12]),
				DaysOfRain:      util.ConvertStringToFloat(row[13]),
				DaysOfMix:       util.ConvertStringToFloat(row[14]),
				DaysOfSnow:      util.ConvertStringToFloat(row[15]),
				MonthlyRainfall: util.ConvertStringToFloat(row[16]),
				MonthlySnowfall: util.ConvertStringToFloat(row[17]),
				InchesPerRain:   util.ConvertStringToFloat(row[18]),
				InchesPerSnow:   util.ConvertStringToFloat(row[19]),
				WindSpeedMin:    util.ConvertStringToFloat(row[20]),
				WindSpeedAvg:    util.ConvertStringToFloat(row[21]),
				WindSpeedMax:    util.ConvertStringToFloat(row[22]),
			}

			regions = append(regions, region)
		}
	}
	return regions
}
