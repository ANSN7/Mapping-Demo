package main

import (
	"fmt"

	"github.com/mashingan/smapping"
)

type Source1 struct {
	Label   string `json:"label"`
	Info    string `json:"info"`
	Version int    `json:"version"`
	Nest nest `json:"nest"`
	Arr   []arr1
}

type arr1 struct {
	Quit string `json:"quit"`
}

type nest struct {
	Base string `json:"base"`
}

type HereticSink struct {
	NahLabel string `json:"label"`
	HahaInfo string `json:"info"`
	Version  string `json:"base"`
	Arr   []arr2
}

type arr2 struct {
	Amber string `json:"quit"`
}

func main() {
	source := Source1{
		Label:   "source",
		Info:    "the origin",
		Version: 1,
		Nest: nest{Base: "we"},
		Arr:      []arr1{{Quit: "hi"}, {Quit: "erw"}},
	}

	// maptags := smapping.MapTags(source.Nest, "json")
	// fmt.Println("maptags:", maptags)

	// hereticsink := HereticSink{}
	// err := smapping.FillStructByTags(&hereticsink, maptags, "json")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("heretic sink:", hereticsink)
	// hereticsink := HereticSink{}

	hereticsink := HereticSink{
		Arr: make([]arr2, len(source.Arr)),
	}

	for i := range source.Arr{
		maptags := smapping.MapTags(source.Arr[i], "json")
		fmt.Println("maptags:", maptags)
		err := smapping.FillStructByTags(&hereticsink.Arr[i], maptags, "json")
		if err != nil {
			panic(err)
		}
		fmt.Println("heretic sink:", hereticsink)
	}
	

}