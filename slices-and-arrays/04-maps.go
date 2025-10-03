// 25 - Maps
package slicesandarrays

import "fmt"

func Maps() {
	cidade := make(map[string]int)

	cidade["SP"] = 9000
	cidade["RJ"] = 800
	cidade["SC"] = 700

	valor, foiEncontrado := cidade["RJ"]
	if foiEncontrado {
		fmt.Println(valor)
	} else {
		fmt.Println("Chave nao existe")
	}
	fmt.Println(cidade)
}
