package fluxodecontrole

import "fmt"

func SalarioIf() {
	salario := 850.00
	var salarioMaisOBonus float64

	salarioMaisOBonus = salario

	if salario < 1000 {
		salarioMaisOBonus = (salarioMaisOBonus + 100)
	}

	fmt.Println("Salario: ", salarioMaisOBonus)
}
