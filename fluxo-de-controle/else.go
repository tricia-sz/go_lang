package fluxodecontrole

import "fmt"

func HabilitacaoElse() {
	idade := 25
	temCarteira := true

	if idade > 18 && temCarteira {
		fmt.Println("Pode dirigir")

	}
}
