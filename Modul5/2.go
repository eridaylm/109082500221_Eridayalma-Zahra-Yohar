package main

import "fmt"

func cetakbintang(k int) {
	if k == 0 {
		return
	}
	fmt.Print("*")
	cetakbintang(k - 1)
}

func pola(barissekarang, ttlbaris int) {
	if barissekarang > ttlbaris {
		return
	}
	cetakbintang(barissekarang)
	fmt.Println()
	pola(barissekarang+1, ttlbaris)
}

func main() {
	var n int
	fmt.Print("Masukkan N : ")
	fmt.Scan(&n)
	pola(1, n)
}