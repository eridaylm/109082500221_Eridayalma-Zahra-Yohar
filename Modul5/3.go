package main

import "fmt"

func tampilfaktor(calon, bilangan int) {
	if calon > bilangan {
		return
	}
	if bilangan%calon == 0 {
		fmt.Printf("%d ", calon)
	}
	tampilfaktor(calon+1, bilangan)
}

func main() {
	var n int
	fmt.Print("Masukkan N : ")
	fmt.Scan(&n)
	fmt.Printf("Faktor dari %d : ", n)
	tampilfaktor(1, n)
	fmt.Println()
}