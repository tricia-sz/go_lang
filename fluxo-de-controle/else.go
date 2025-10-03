package fluxodecontrole

import "fmt"

func SalarioElse() {
	salario := 850.00
	var salarioMaisOBonus float64

	salarioMaisOBonus = salario

	if salario < 1000 {
		salarioMaisOBonus = (salarioMaisOBonus + 100)
	} else {
		fmt.Println("Salario: ", salarioMaisOBonus)

	}

	fmt.Println("Salario: ", salarioMaisOBonus)
}
