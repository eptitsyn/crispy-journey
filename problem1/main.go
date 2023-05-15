package main

import "fmt"

/*
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
//done

type Human struct {
	name string
	age  uint
}

func (h *Human) Yell() {
	fmt.Printf("Hello, world, my name is %s\n", h.name)
}

type Action struct {
	Human
}

func (h *Action) Yell() {
	fmt.Printf("Action!!")
	h.Human.Yell()
}

func main() {
	action := Action{Human{
		name: "Bob",
		age:  21,
	}}
	action.Yell()
}
