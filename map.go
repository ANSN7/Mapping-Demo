package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  age
}

type age struct {
	mine int
}

var b = []byte(`{"Persons":[                                                                       
                   {"Name": "John", "Age" : {"mine": 12} }                                                 
               ]}`)

func main() {
	var f interface{}
	err := json.Unmarshal(b, &f)
	// fmt.Println(b)
	fmt.Println(f)
	if err != nil {
		panic("OMG!")
	}

	m := f.(map[string]interface{})
	fmt.Println(m)

	for k, v := range m {
		switch val := v.(type) {
		case string:
			fmt.Println(k, "is string", val)
		case float64:
			fmt.Println(k, "is float64", val)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for index, value := range val {
				fmt.Println(index, value)
			}
		default:
			fmt.Println(k, "is of unknown type")
		}
	}

	n := m["Persons"].([]interface{})
	fmt.Println(n)

	// persons := make([]*Person, len(n))

	// for i := range n {
	// 	name := n[i].(map[string]interface{})["Name"].(string)
	// 	age := n[i].(map[string]interface{})["Age"].(float64)
	// 	persons[i] = &Person{name, int(age)}
	// 	fmt.Println(name, int(age))
	// }
}
