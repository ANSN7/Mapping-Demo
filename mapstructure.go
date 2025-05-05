package main

import (
	"fmt"
	// "encoding/json"
	"github.com/mitchellh/mapstructure"
)

type Person1 struct {
	Name string `mapstructure:"person_name"`
	Age  age1   `mapstructure:"person_age,squash"`
}

type age1 struct {
	base string `mapstructure:"base_p"`
}

func main() {

	input := map[string]interface{}{
		"person_name": "Mitchell",
		"person_age":  age1{base: "tyu"},
		// "we":0,
	}

	var result Person1
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

    input = map[string]interface{}{
		"base_p": "tyu",
	}

	var result2 age1
	err = mapstructure.Decode(input, &result2)

	// fmt.Printf("%#v", result)
	fmt.Println(result)
	fmt.Println(result2)
	// fmt.Printf("%#v", result2)
}
