package main

import "fmt"

func main() {

/* 	for i := 0; i < 5; i++ {
		fmt.Println(i)
	} */

	for j := 4; j >= 0; j-- {
		if j == 3	{
			continue // skip or ignore this sentence, instance or condition iteration and continue with the next iteration
		}
		fmt.Println(j)
	}

/* 	for j := 4; j >= 0; j-- {
		if j == 2	{
			break // Rompe o termina las iteraciones cuando se cumple la condicion
		}
		fmt.Println(j)
	} */

	matriz := [3][3]int{}
	valor := 0

	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			valor++
			matriz[k][l] = valor
		}
	}

	fmt.Println(matriz)

	for fila := 0; fila < 3; fila++ {
		fmt.Println(matriz[fila])
		/* for col := 0; col < 3; col++ {
			fmt.Println(matriz[fila][col])
		} */
	}
}