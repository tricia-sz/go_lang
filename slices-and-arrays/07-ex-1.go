// EX 01
package slicesandarrays

import "fmt"

/*
 Criar um array com 2 posiçoes de inteiros e armazenar em uma variavel a soma total da lista;
A variavel deve ser imprimida no console.
*/

func Ex1() {
	lista := [2]int{1, 2}
	fmt.Println(lista)
	soma := lista[0] + lista[1]
	fmt.Println(soma)

}

/*
outline:
[1 2]
3

*/
