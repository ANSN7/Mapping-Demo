package main

import (
	"fmt"

	res "k8s.io/apimachinery/pkg/api/resource"
)

type CodecoAppResourceBase struct {
	//+kubebuilder:validation:default=100
	Cpu *res.Quantity `json:"cpu,omitempty"`

	//+kubebuilder:validation:default=8
	Mem *res.Quantity `json:"mem,omitempty"`
}

func main() {
	// memorySize := res.MustParse("5Gi")
	// fmt.Print(memorySize.Value())

	cpu := res.MustParse("100m")
	mem := res.MustParse("5")

	codeco := CodecoAppResourceBase{Cpu: &cpu, Mem: &mem}
	fmt.Println(codeco)

	// cpu, err := res.ParseQuantity("100m")
	// if err != nil {
	// 	fmt.Print("error")
	// }
	// fmt.Print(cpu)
}
