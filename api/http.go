package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"weather/util"

	"github.com/gorilla/mux"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	allRecords := util.CsvToJson()
	jsonData, err := json.MarshalIndent(allRecords, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(jsonData))

	// err = os.WriteFile("out.json", jsonData, 0655)
	// if err != nil {
	// 	fmt.Println(err)
	//}
}

func filerByQuery(w http.ResponseWriter, r *http.Request) {
	var date, weather, limit string
	values := r.URL.Query()
	for k, v := range values {
		if k == "date" {
			date = strings.Join(v, "")
		}
		if k == "weather" {
			weather = strings.Join(v, "")
		}
		if k == "limit" {
			limit = strings.Join(v, "")
		}
	}
	fmt.Println(date, weather, limit)
	// date := r.URL.Query().Get("date")
	// weather := r.URL.Query().Get("weather")
	// limit := r.URL.Query().Get("limit")

	allRecords := util.CsvToJson()
	var lists, datelist, weatherlist []util.Weather

	if date != "" {
		for _, record := range allRecords {
			if record.Date == date {
				list := util.Weather{
					Date:          record.Date,
					Precipitation: record.Precipitation,
					Temp_max:      record.Temp_max,
					Temp_min:      record.Temp_min,
					Wind:          record.Wind,
					Weather:       record.Weather,
				}
				datelist = append(datelist, list)
			}
		}
	}

	if weather != "" {
		for _, record := range allRecords {
			if record.Weather == weather {
				list := util.Weather{
					Date:          record.Date,
					Precipitation: record.Precipitation,
					Temp_max:      record.Temp_max,
					Temp_min:      record.Temp_min,
					Wind:          record.Wind,
					Weather:       record.Weather,
				}
				weatherlist = append(weatherlist, list)
			}
		}
	}

	if (date != "" && limit != "") || (weather != "" && limit != "") {
		if len(datelist) == 0 {
			lists = weatherlist
			lists = ifLimit(limit, lists)
		} else {
			lists = datelist
			lists = ifLimit(limit, lists)
		}

	} else if date != "" && weather != "" {
		for _, data := range datelist {
			if data.Weather == weather {
				out := util.Weather{
					Date:          data.Date,
					Precipitation: data.Precipitation,
					Temp_max:      data.Temp_max,
					Temp_min:      data.Temp_min,
					Wind:          data.Wind,
					Weather:       data.Weather,
				}
				lists = append(lists, out)
			} else {
				break
			}

		}

	} else if limit != "" {
		lists = ifLimit(limit, allRecords)
	}

	output, err := json.MarshalIndent(lists, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(output))
}

func ifLimit(limit string, data []util.Weather) []util.Weather {
	i, err := strconv.Atoi(limit)
	if err != nil {
		fmt.Println(err)
	}

	output := data[:i]
	return output
}

func HandleRequests() {
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/", handleGet).Methods("GET")
	route.HandleFunc("/query", filerByQuery).Methods("GET")
	http.ListenAndServe(":8080", route)

}
