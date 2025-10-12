package structs

import (
	"fmt"
	model "learngo/model"
)

func heranca() {
	fmt.Println("Iniciando...")
	automovelMoto := model.Automovel{
		Ano:    2022,
		Placa:  "TRX-2020",
		Modelo: "BMW",
	}

	moto := model.Moto{
		Automovel:   automovelMoto,
		Cilindradas: 125,
	}

	fmt.Println(moto.Modelo)
}
