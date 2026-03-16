package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	hasil1 := f(g(h(a)))
	fmt.Println(hasil1)

	hasil2 := g(h(f(b)))
	fmt.Println(hasil2)

	hasil3 := h(f(g(c)))
	fmt.Println(hasil3)
}