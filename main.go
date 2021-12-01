package main

import "fmt"

func main(){
	declaracionCorta()
}

func saludoPersona(nombrePersona string, edadPersona int){
	fmt.Println("Hola " + nombrePersona)
	fmt.Printf("%d", edadPersona)
}

var variableGlobal string = "Soy una variable global"

func declaracionCorta(){
	// ...interface{} -> varios parámetros de cualquier tipo
	fmt.Println("Cadena", 12, true)
	edadPersona := 25
	fmt.Println(edadPersona)

	var nombreEstudiante string = "Soy el estudiante"
	fmt.Println(nombreEstudiante)
	fmt.Println(variableGlobal)
}
