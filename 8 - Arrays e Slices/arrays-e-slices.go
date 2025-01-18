package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	fmt.Println(array1)

	array1[0] = 29
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]string{"João", "Paulo", "Henrique", "José"}
	fmt.Println(array3)

	slice := []string{"Giovanna","Adriana","Iracema","Rosa","Maria","Aparecida"}
	fmt.Println(slice)
	fmt.Println(len(slice)) // tamanho
	fmt.Println(cap(slice)) // capacidade
	slice = append(slice, "Luciana")
	fmt.Println(len(slice)) // tamanho
	fmt.Println(cap(slice)) // capacidade
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[2] = "Posição Nova"
	fmt.Println(slice2)


	// Arrays Internos
	fmt.Println("--------------------------------------")
	slice3 := make([]float32, 5, 6)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade
	fmt.Println(slice3)
	slice3 = append(slice3, 0)
	slice3 = append(slice3, 0)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade
	fmt.Print(slice3)
}