package main

import "fmt"

func main() {

	mes := "octubre"

	dia := 23

	fmt.Println(dia, mes)

	fmt.Printf(" Hoy es %d de %s \n", dia, mes)

	fmt.Printf(" Hoy es %v de %v \n", dia, mes)

	fmt.Printf("La variable mes es un %T, y la variable dia es un %T \n", mes, dia)

	fecha := fmt.Sprintf(" Hoy es %v de %v \n", dia, mes)

	fmt.Println(fecha)

	// DESAFIO

	// Guardar en una variable "mascota" el nombre de tu mascota,
	// Guardar en la variable "edad" la edad de tu mascota
	// Guardar en la variable variable "esPerro" true si es un perro, sino es un false.
	// Guardar en la variable "miMascota" un mensaje utilizando las variables creadas arriba.
	// Imprimir el mensaje en pantalla
	// Imprimir de que tipo son las variables creadas arriba

	mascota := "Trufa"

	edad := 1

	esPerro := false

	miMascota := fmt.Sprintf("Mi mascota se llama %s, tiene %d año y es %v que es un perro ", mascota, edad, esPerro)

	fmt.Println(miMascota)

	fmt.Printf("mascota es un %T, edad es un %T y esPerro es un %T \n", mascota, edad, esPerro)

}
