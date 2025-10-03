package learngo

import "fmt"

func CriandoTipo() {
	type hotdog int

	var b hotdog
	var c int

	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}

/*
	InputTerminal:
	main.hotdog
	int
*/
