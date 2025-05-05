package main

import (
	"encoding/json"
	"fmt"
	"github.com/ANSN7/mapper"
)

type ModelA struct {
	PrimerNivel  string `json:"primer_nivel" mapper:"nivelA"`
	SegundoNivel string `json:"segundo_nivel,omitempty" mapper:"nodo.nivelB"`
	TercerNivel  string `json:"tercer_nivel" mapper:"nodo_dos.sub_nodo.nivelC"`
	// ObjRecursivo ModelA   `json:"obj_recursivo" mapper:"recursivo"`
	Array []work1 `json:"arr_obj" mapper:"array"`
	Email string  `json:"email" mapper:"nodo_dos.sub_nodo.email_validado"`
}

type ModelB struct {
	NivelA string `json:"nivelA"`
	Nodo   struct {
		NivelB string `json:"nivelB"`
	} `json:"nodo"`
	NodoDos struct {
		SubNodo struct {
			NivelC        string `json:"nivelC"`
			EmailValidado string `json:"email_validado"`
		} `json:"sub_nodo"`
	} `json:"nodo_dos"`
	// Recursivo ModelB   `json:"recursivo"`
	Array []work2 `json:"array"`
}

type work1 struct {
	Base string `mapper:"war"`
	No   struct {
		Niv string `json:"nivelBB" mapper:"ss"`
	} `json:"nod"`
}

type work2 struct {
	War string `json:"war"`
	S string `json:"ss,omitempty"`
	
}

func main() {
	src := ModelA{PrimerNivel: "as", SegundoNivel: "er", TercerNivel: "sc", Email: "wer", Array: []work1{{Base: "hi", No: struct{Niv string "json:\"nivelBB\" mapper:\"ss\""}{Niv: "cdjc"}},}}
	dest := d1(src)
	fmt.Println(dest)
}

func d1(source ModelA) ModelB {
	fmt.Print("..........................")
	target := ModelB{}
	errMessage := mapper.Apply(source, &target)
	//El mapper se va a aplicar a todos los campos que se pueda
	//aquellos que fallen quedaran seran retornados en una cadena de errores por cada falla.
	//Por lo que no es condición necesaria que todos los campos se mapeen
	//Esto para evitar romper por algun tipo de dato
	//En tal caso se puede loguear el error y monitorear
	if errMessage != nil {
		//log errMessage
		fmt.Println("errMessage", errMessage)
	}

	jsonData, err := json.Marshal(&target)
	if err != nil {
		fmt.Println("Error marshaling ModelB:", err)
	} else {
		jsonString := string(jsonData)
		fmt.Println(jsonString)
	}

	return target
}
