package util

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Weather struct {
	Date          string `json:"date"`
	Precipitation string
	Temp_max      string
	Temp_min      string
	Wind          string
	Weather       string
}

type WeatherReport []Weather

func CsvToJson() WeatherReport {
	filepath := "/weather/seattleweather.csv"
	csvfile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	var records []Weather

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		record := Weather{
			Date:          row[0],
			Precipitation: row[1],
			Temp_max:      row[2],
			Temp_min:      row[3],
			Wind:          row[4],
			Weather:       row[5],
		}
		records = append(records, record)
	}
	return records

}
