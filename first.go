package main

import "fmt"

var name string

func main() {
		name = "Guilherme"
		x := showMyName()
    fmt.Println(x)
}

func showMyName() string {
	 return name
}