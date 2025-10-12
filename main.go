package main

import "fmt"

func main() {
	x := 5
	y := &x

	*y = 10

	fmt.Println("___FUNC MAIN____")

	fmt.Println(x, *y)
	fmt.Println(&x, y)

	ImprimirValores(&x, y)

}

func ImprimirValores(x *int, y *int) {
	fmt.Println("___FUNC IMPRIME VALORES_____")

	fmt.Println(x, y)
	fmt.Println(&x, &y)

}
