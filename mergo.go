package main

import (
	"fmt"
	"dario.cat/mergo"
)

type Foo struct {
	A string
	B int64
	C base
}

type Foo1 struct {
	A string
}

type base struct{
	arm1 string
}

func main() {
	src := Foo{
		A: "one",
		B: 2,
		C: base{arm1: "we"},
	}
	dest := Foo1{
		A: "two",
	}
	mergo.Merge(&dest, src)
	fmt.Println(dest)
	// Will print
	// {two 2}
}