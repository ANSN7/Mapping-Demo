package main

import (
    "encoding/json"
    "fmt"
	"github.com/tidwall/pretty"
)

// Define a struct to represent your JSON data
type MyData struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Country string `json:"country"`
}

func main() {
    // Create an instance of the struct and populate its fields
    data := MyData{
        Name:    "John Doe",
        Age:     30,
        Country: "USA",
    }

    // Marshal the struct into JSON
    jsonData, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return
    }

    // Print the JSON data
    fmt.Println(string(jsonData))

	result := pretty.Pretty(jsonData)
    fmt.Println(string(result))

}
