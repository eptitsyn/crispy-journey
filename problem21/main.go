package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

type Animal interface {
	Eat()
}
type Dog struct {
}

func (d *Dog) Eat() {
	fmt.Println("Dog eats")
}

type Cat struct {
}

func (c *Cat) EatsBestFood() {
	fmt.Println("Cat eats")
}

type CatAdapter struct {
	Cat *Cat
}

func (a *CatAdapter) Eat() {
	a.Cat.EatsBestFood()
}

type PetFood struct{}

func (f *PetFood) Feed(pet Animal) {
	fmt.Println("feeding pet")
	pet.Eat()
}

func main() {
	food := &PetFood{}
	cat := &CatAdapter{Cat: &Cat{}}
	dog := &Dog{}

	food.Feed(dog)
	food.Feed(cat)
	//food.Feed(Cat{}) - err
}
