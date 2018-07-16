package main

import "fmt"

func constantes() {
	const nombre = "Carlos"

	fmt.Println(nombre)
}

func tipoDatos() {
	var a int //int es el alias del tipo int32 o int64 dependiendo del sistema donde se va a compilar
	var b int8

	a = 12569
	b = 32
	//c := a + b genera error x diferencia entre los tipos de datos

	/**
	* Aplicamos un Casting: El casting es el cambio temporal de un tipo de datos para realizar una operacion
	 */
	c := a + int(b)

	fmt.Println(c)
	/*Aplicando formato a la salida*/
	n := "Pedro"
	fmt.Printf("Hola %s como estas \n", n)
	fmt.Printf("Este es el numero %d \n", c)
	fmt.Printf("El tipo de dato de c es: %T", c)
}

func main() {
	constantes()
	tipoDatos()
}
