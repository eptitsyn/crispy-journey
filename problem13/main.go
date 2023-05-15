package main

/*
Поменять местами два числа без создания временной переменной.
*/
func main() {
	a := 1
	b := 2
	println(a, b)
	a, b = b, a
	println(a, b)
}
