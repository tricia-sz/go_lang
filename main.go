package main

import "fmt"

func main() {
	x := 5
	y := &x

	*y = 10

	fmt.Println(x, *y)
	fmt.Println(&x, y)

	ImprimirValores(&x, y)

}

func ImprimirValores(x *int, y *int) {
	*x = 20
}
