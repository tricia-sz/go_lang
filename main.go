package main

import (
	"fmt"
	"os"
)

func showText() {
	fmt.Println("Finalizando a manipulacao do arquivo")
}

func main() {
	file, err := os.Create("./settings.txt")
	defer file.Close()
	defer showText()

	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte("teste"))
	if err != nil {
		panic(err)
	}
}
