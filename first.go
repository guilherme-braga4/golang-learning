package main

import "fmt"

type Teste int 

var x Teste

func main() {
		x := 1

		fmt.Println(x)
		fmt.Println("%T\n", x)
		x = 42
		fmt.Println(x)
	}
