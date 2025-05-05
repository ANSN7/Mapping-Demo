package main
import (
    dc "github.com/fluidtruck/deepcopy"
	"fmt"
)

type Struct0 struct {
    Base0 string
	Conditions []AC0
}

type Struct1 struct {
	Base1 string
	Conditions []AC1
	Node map[string]float64
}

type AC0 struct {
	Reason string 
	H00 string
	Entity string
}

type AC1 struct {
	Reason string  
	H11 string
	Entity string
}



func main() {
    a := Struct0{Base0: "we", Conditions: []AC0{
		{Reason: "reason1", H00: "H00A", Entity: "en"}, 
		{Reason: "reason2", H00: "H00AA", Entity: "enen"}, 
	}}
    b := Struct1{
		Base1: "weeeee", 
		Node: map[string]float64{
			"c1": 0.8,
			"c2": 0.2,
		},
	}
    err := dc.DeepCopy(a, &b)
    if err != nil {
        fmt.Println(err)
    }
    
    fmt.Println(b)
    fmt.Println(b)
}