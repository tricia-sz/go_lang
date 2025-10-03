// 26 - RANGE

package slicesandarrays

import "fmt"

func Range() {
	cidade := make(map[string]int)

	cidade["SP"] = 9000
	cidade["RJ"] = 800
	cidade["SC"] = 700

	for chave, valor := range cidade {
		fmt.Println("Cidade:", chave, "Habitantes:", valor)
	}

}
