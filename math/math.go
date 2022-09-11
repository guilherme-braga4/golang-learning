package math

import "fmt"

func Sum (num1, num2 int) int {
   fmt.Printf("Acabei de Receber 2 variáveis: %v e %v \n", num1, num2)
	 fmt.Println("Agora vou somar esses dois")
	 resultado := num1 + num2
	 return resultado
}

func Subtraction (num1, num2 int) int {
	fmt.Println("Acabei de Receber 2 variáveis: %v e %v \n", num1, num2)
	fmt.Println("Agora vou subtrair esses dois")
	resultado := num1 - num2
	return resultado
}

func Results (num1, num2 int) (int, int) { 
	return Subtraction(num1, num2), Sum(num1, num2)
}