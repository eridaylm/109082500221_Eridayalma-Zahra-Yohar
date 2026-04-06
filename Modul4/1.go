package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 2; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int) int {
	var fn, fnr int
	faktorial(n, &fn)
	faktorial(n-r, &fnr)
	return fn / fnr
}

func kombinasi(n, r int) int {
	var fn, fr, fnr int
	faktorial(n, &fn)
	faktorial(r, &fr)
	faktorial(n-r, &fnr)
	return fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	fmt.Print("Masukkan a b c d : ")
	fmt.Scan(&a, &b, &c, &d)

	p1 := permutasi(a, c)
	c1 := kombinasi(a, c)
	p2 := permutasi(b, d)
	c2 := kombinasi(b, d)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}