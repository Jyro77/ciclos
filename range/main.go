package main

import "fmt"

func main() {

	// para hacer iteraciones

	num := []string{"0", "1", "2", "3", "4", "5", "6"}

	// range devuelve i = index y v = value, se puede omitir, o dejar de recibir (evitando la declaracion) del valor, pero no la del index
	for i, v := range num {
		fmt.Println(i, v)
	}
	paises := map[string]string{"RD": "Republica Dominicana", "CO": "Colombia", "BR": "Brasil", "PR": "Puerto Rico"}
	// tambien se puede ignorar una variable (valor recibido) al declararla con _ ejemplo:
	for _, val := range paises {
		fmt.Println(val)
	}

	// se puede iterar strings
	frase := "Hola Mundo!"
	
	for _, val := range frase {
		fmt.Println(val)
		// pero printará el valor de la letra segun bit... para convertir eso a letra se debe de pasar la variable en la declaracion por el convertidor. Ejemplo: "string(val)"
	}

	//tambien se puede hacer la declaracion de lo que se va a iterar dentro de la misma declaracion del for
	for _,v := range []int{1,2,3,4,5}{
		fmt.Println("valor es", v)
	}

	for _,v := range "Curso de GO desde cero" {
		fmt.Println(string(v))
	}
}