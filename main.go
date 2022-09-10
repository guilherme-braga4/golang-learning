package main

import (
	"modulo/children"
	"fmt"
  "github.com/badoux/checkmail"
)

type Teste int 

var x Teste

func main() {
		x := 1

		fmt.Println(x)
		fmt.Println("%T\n", x)
		x = 42
		fmt.Println(x)
		children.Filho()
		testeEmail := checkmail.ValidateFormat("gui@gmail.com")
		fmt.Println(testeEmail)
	}
