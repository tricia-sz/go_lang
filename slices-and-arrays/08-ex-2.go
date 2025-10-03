// EX 02

package slicesandarrays

import "fmt"

/*
Dado um slice com os itens "2, 8, 3, 10,5, 4, 7, 9, 1" que vao de 1 a 10.
 Efetuar a soma de duas variaveis, a primeira numeros de 1 a 5 e a segunda de 6 a 10
imprimir os dois resultados
*/

func Ex2() {
	lista := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	numeroAte5 := 0
	numeroAte10 := 0

	for i := 0; i < len(lista); i++ {
		if lista[i] <= 5 {
			numeroAte5 = numeroAte5 + lista[i]
		} else {
			numeroAte10 += lista[i]
		}

	}

	fmt.Println(numeroAte5)
	fmt.Println(numeroAte10)
}
