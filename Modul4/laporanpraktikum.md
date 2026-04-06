# <h1 align="center">Laporan Praktikum Modul 4 - Algoritma Pemrograman 2 </h1>
<p align="center">Eridayalma Zahra Yohar - 109082500221</p>

## Modul 4 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)
#### 1.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul4/Output/output-soal1.png)
Penjelasan : 
Program ini ngebaca 4 bilangan a, b, c, d terus ngitung permutasi dan kombinasi dari (a,c) dan (b,d). Program punya prosedur faktorial yang ngitung nilai faktorial pake perulangan. Fungsi permutasi dan kombinasi manggil prosedur faktorial terus menghitung sesuai rumus P(n,r)=n!/(n-r)! dan C(n,r)=n!/(r!×(n-r)!).

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan totalsoal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.
#### 2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul4/Output/output-soal2.png)
Penjelasan : 
Program ini nentuin pemenang lomba coding dari daftar peserta. Setiap peserta punya nama dan 8 nilai waktu (dalam menit). Waktu 301 menit atau lebih dianggap gagal. Program pake prosedur hitungskor yang ngebaca 8 waktu, ngitung jumlah soal berhasil (waktu < 301) dan total waktu dari soal berhasil. Program ngebaca nama dan manggil prosedur berulang sampe menemukan kata "Selesai". Pemenang dipilih itu berdasarkan jumlah soal paling banyak, dan kalo sama, berdasarkan total waktu terkecil. Hasil nya munculin nama pemenang, jumlah soal, dan total waktu.