// 27 - Delete

package slicesandarrays

import "fmt"

func Delete() {
	cidade := make(map[string]int)

	cidade["SP"] = 9000
	cidade["RJ"] = 800
	cidade["SC"] = 700
	cidade["BH"] = 500

	delete(cidade, "SC")

	for chave, valor := range cidade {
		fmt.Println("Cidade:", chave, "Habitantes:", valor)
	}

}
