package main

import "fmt"

func main(){
  /*declaracion de variable 1*/
  var nombre1, nombre2 string
  nombre1 = "Carlos"
  nombre2 = "Eduardo"

  /*declaracion de variable 2*/
  apellido1, apellido2 := "Villanueva", "Altuna"

  /*Reasignar un nuevo valor a una variable ya creada. Si usamos una variable existente con el operador := debemos
    crear al menos una nueva variable al lado de lo contrario mostrara un error.*/
  nombre1, edad := "Juan", 39

  fmt.Println(nombre1, nombre2, apellido1, apellido2, edad)
}
