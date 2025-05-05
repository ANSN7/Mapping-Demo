package main

import (
	"fmt"
	"reflect"

	"github.com/m7shapan/njson"
)

var jsonString string = `
	{
		"coord": {
			"lon": -0.13,
			"lat": 51.51
		},
		"weather": [
			{
				"id": 300,
				"main": "Drizzle",
				"description": "light intensity drizzle",
				"icon": "09d"
			}
		],
		"base": "stations",
		"main": {
			"temp": 280.32,
			"pressure": 1012,
			"humidity": 81,
			"temp_min": 279.15,
			"temp_max": 281.15
		},
		"visibility": 10000,
		"wind": {
			"speed": 4.1,
			"deg": 80
		},
		"clouds": {
			"all": 90
		},
		"dt": 1485789600,
		"sys": {
			"type": 1,
			"id": 5091,
			"message": 0.0103,
			"country": "GB",
			"sunrise": 1485762037,
			"sunset": 1485794875
		},
		"id": 2643743,
		"name": {"ew": "London", "ss": 1485794875},
		"cod": 200
	}
`

func main() {
	njsonUnmarshaling()
}


func njsonUnmarshaling() {

	type Weather struct {
		Location       string  `njson:"name"`
		Weather        string  `njson:"weather.0.main"`
		Description    string  `njson:"weather.0.description"`
		Temperature    float32 `njson:"main.temp"`
		MinTemperature float32 `njson:"main.temp_min"`
		MaxTemperature float32 `njson:"main.temp_max"`
	}

	var weather = Weather{}
	value := reflect.ValueOf(weather)
	fmt.Println(value)
	err := njson.Unmarshal([]byte(jsonString), &weather)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", weather)

}


// func PrintFields(codeco codecov1alpha1.CodecoApp) {
// 	val := reflect.ValueOf(codeco)
// 	for i := 0; i < val.Type().NumField(); i++ {
// 		fmt.Println(val.Type().Field(i).Tag.Get("mapper"))
// 	}
// }