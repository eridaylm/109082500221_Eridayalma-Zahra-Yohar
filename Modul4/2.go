package main

import "fmt"

func hitungskor(soal *int, skor *int) {
	*soal = 0
	*skor = 0
	var waktu int
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor int
	maxSoal := -1
	minSkor := 0
	adapeserta := false

	fmt.Println("Masukkan data peserta (akhiri dengan \"Selesai\"):")
	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}
		hitungskor(&soal, &skor)
		adapeserta = true

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		}
	}

	if adapeserta {
		fmt.Println(pemenang, maxSoal, minSkor)
	} else {
		fmt.Println("Tidak ada peserta")
	}
}