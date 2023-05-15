package main

import "fmt"

/*
	Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool,
	channel из переменной типа interface{}.
*/
//done

func determineType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Variable is of type int")
	case string:
		fmt.Println("Variable is of type string")
	case bool:
		fmt.Println("Variable is of type bool")
	case chan int:
		fmt.Println("Variable is of type channel of type int")
	default:
		fmt.Println("Variable type is unknown")
	}
}

func main() {
	var a interface{}
	a = 42
	determineType(a)
	a = "Hello, world!"
	determineType(a)
	a = true
	determineType(a)
	ch := make(chan int)
	a = ch
	determineType(a)
	a = 1.23
	determineType(a)
}
