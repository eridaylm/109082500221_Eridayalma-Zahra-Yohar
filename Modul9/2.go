package main

import (
	"fmt"
	"math"
)

type tabelInt [100]int

func tampilSemua(t tabelInt, n int) {
	fmt.Print("Semua elemen : ")
	for i := 0; i < n; i++ {
		fmt.Print(t[i], " ")
	}
	fmt.Println()
}

func tampilGanjil(t tabelInt, n int) {
	fmt.Print("Indeks ganjil : ")
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(t[i], " ")
		}
	}
	fmt.Println()
}

func tampilGenap(t tabelInt, n int) {
	fmt.Print("Indeks genap : ")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(t[i], " ")
		}
	}
	fmt.Println()
}

func tampilKelipatanX(t tabelInt, n int, x int) {
	fmt.Print("Indeks kelipatan ", x, " : ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(t[i], " ")
		}
	}
	fmt.Println()
}

func hapusIndex(t *tabelInt, n *int, idx int) {
	for i := idx; i < *n-1; i++ {
		t[i] = t[i+1]
	}
	*n = *n - 1
}

func hitungRata(t tabelInt, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total = total + t[i]
	}
	rata := float64(total) / float64(n)
	return rata
}

func hitungStdDev(t tabelInt, n int) float64 {
	rata := hitungRata(t, n)
	jumlah := 0.0
	for i := 0; i < n; i++ {
		selisih := float64(t[i]) - rata
		jumlah = jumlah + (selisih * selisih)
	}
	hasil := math.Sqrt(jumlah / float64(n))
	return hasil
}

func hitungFrekuensi(t tabelInt, n int, angka int) int {
	count := 0
	for i := 0; i < n; i++ {
		if t[i] == angka {
			count++
		}
	}
	return count
}

func main() {
	var tab tabelInt
	var n int

	fmt.Print("Masukkan jumlah elemen : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Elemen ke-", i, " : ")
		fmt.Scan(&tab[i])
	}

	fmt.Println()

	// a
	tampilSemua(tab, n)

	// b
	tampilGanjil(tab, n)

	// c
	tampilGenap(tab, n)

	// d
	var x int
	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)
	tampilKelipatanX(tab, n, x)

	// e
	var idx int
	fmt.Print("Masukkan indeks yang dihapus : ")
	fmt.Scan(&idx)
	hapusIndex(&tab, &n, idx)
	fmt.Print("Array setelah dihapus : ")
	tampilSemua(tab, n)

	// f
	fmt.Printf("Rata-rata : %.2f\n", hitungRata(tab, n))

	// g
	fmt.Printf("Standar deviasi : %.2f\n", hitungStdDev(tab, n))

	// h
	var cari int
	fmt.Print("Cari frekuensi angka : ")
	fmt.Scan(&cari)
	fmt.Println("Frekuensi", cari, ":", hitungFrekuensi(tab, n, cari), "kali")
}