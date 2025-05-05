package main

import (
	"log"

	"github.com/rookiecj/go-langext/mapper"
)

func main() {
	

	type align struct {
		Base string 
	}

	
	type common struct {
		Desc string
		A align 
	}

	type src struct {
		Seq   int  `mapper:"ID"`
		Label string `mapper:"Name"`
		Addr  string
		common
	}

	type dest struct {
		ID   int   
		Name string 
		Addr string
		common
	}

	srcValue := src{
		Seq:   13,
		Label: "Prime",
		Addr:  "Number",
		common: common{
			Desc: "Embedded",
			A: align{Base: "nvkjf"},
		},
	}

	destValue := dest{}

	m := mapper.NewMapperWithTag("mapper")

	if err := m.Map(&destValue, srcValue); err != nil {
		panic(err)
	}
	log.Printf("srcValue: %v", srcValue)
	log.Printf("destValue: %v", destValue)
}