# <h1 align="center">Laporan Praktikum Modul 9 - Algoritma Pemrograman 2 </h1>
<p align="center">Eridayalma Zahra Yohar - 109082500221</p>

## Modul 9 

### 1.Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥, 𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥, 𝑦) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".


#### 1.go

```go
package main

import (
	"fmt"
	"math"
)

type titik struct {
	x int
	y int
}

type lingkaran struct {
	pusat titik
	r     int
}

func jarak(p titik, q titik) float64 {
	hasil := math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
	return hasil
}

func didalam(c lingkaran, p titik) bool {
	d := jarak(c.pusat, p)
	if d <= float64(c.r) {
		return true
	}
	return false
}

func main() {
	var l1 lingkaran
	var l2 lingkaran
	var p titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&p.x, &p.y)

	cek1 := didalam(l1, p)
	cek2 := didalam(l2, p)

	if cek1 == true && cek2 == true {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if cek1 == true {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if cek2 == true {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul9/Output/output-soal1.png)
Penjelasan : 
Program ini untuk nentuin posisi titik terhadap dua lingkaran. Pertama-tama program nerima input koordinat titik pusat dan radius dari dua lingkaran, terus nerima input koordinat titik yang mau dicek posisinya.
Program pake dua tipe bentukan yaitu titik buat nyimpen koordinat x dan y, terus lingkaran buat nyimpeb titik pusat dan radius. Ada dua fungsi utama yang dipake, yaitu fungsi jarak buat ngitung jarak antara dua titik pake rumus Euclidean √((x1-x2)² + (y1-y2)²), dan fungsi didalam untuk ngecek apakah itu titik berada di dalam lingkaran, caranya bandingin jarak titik ke pusat lingkaran dengan radiusnya. Kalo jarak tersebut lebih kecil atau sama dengan radius, maka titik dinyatakan berada di dalam lingkaran.
Setelah dua pengecekan itu, program bakal nampilin salah satu dari empat kemungkinan output yaitu titik berada di dalam lingkaran 1 dan 2, hanya di dalam lingkaran 1, hanya di dalam lingkaran 2, atau di luar keduanya.

### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:
a. Menampilkan keseluruhan isi dari array.
b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah
genap).
d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh
dari masukan pengguna.
e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.
Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array
tersebut.
h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi
tersebut.

#### 2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul9/Output/output-soal2.png)
Penjelasan : 
Program ini digunakan buat nyimpen sekumpulan bilangan bulat ke array, terus dia ngelakuin banyak operasi ke array tersebut.
Pertama program minta input jumlah elemen yang ingin dimasukkan, terus minta input nilai tiap elemennya satu per satu. Kalao array udah ke isi, program melakukan beberapa hal berikut:
- Menampilkan semua elemen dilakukan dengan looping dari indeks 0 sampai n-1 lalu mencetak tiap elemennya.
- Menampilkan elemen indeks ganjil dan genap dilakukan dengan mengecek sisa bagi indeks terhadap 2. Jika i % 2 != 0 maka indeks tersebut ganjil, sebaliknya jika i % 2 == 0 maka genap.
- Menampilkan elemen indeks kelipatan x dilakukan dengan cara yang sama, yaitu mengecek apakah i % x == 0.
- Menghapus elemen pada indeks tertentu dilakukan dengan cara menggeser semua elemen yang ada di sebelah kanan indeks tersebut satu posisi ke kiri, kemudian nilai n dikurangi 1 sehingga elemen terakhir yang sudah tidak relevan tidak ikut ditampilkan.
- Menghitung rata-rata dilakukan dengan menjumlahkan semua elemen lalu dibagi dengan jumlah elemen n.
- Menghitung standar deviasi dilakukan dengan mencari selisih tiap elemen terhadap rata-rata, mengkuadratkannya, menjumlahkan semua hasil kuadrat tersebut, membaginya dengan n, lalu diakarkan menggunakan fungsi math.Sqrt.
- Menghitung frekuensi dilakukan dengan looping seluruh array dan menghitung berapa kali bilangan yang dicari muncul.

### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.


#### 3.go

```go
package main

import "fmt"

const NMAX int = 100

type tabelHasil [NMAX]string

func main() {
	var klubA string
	var klubB string
	var hasil tabelHasil
	var n int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	nomer := 1
	for {
		var skorA int
		var skorB int
		fmt.Print("Pertandingan ", nomer, " : ")
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[n] = klubA
		} else if skorB > skorA {
			hasil[n] = klubB
		} else {
			hasil[n] = "Draw"
		}

		n++
		nomer++
	}

	for i := 0; i < n; i++ {
		fmt.Println("Hasil", i+1, ":", hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul9/Output/output-soal3.png)
Penjelasan :
Program ini buat ngerekap hasil pertandingan dua klub sepak bola. Program minta input nama kedua klub, terus secara berulang ulang minta input skor dari masing-masing klub tiap pertandingan.
Setiap skor dimasukkan, program langsung nentuin siapa yang menang. Kalo skor klub A lebih besar maka nama klub A disimpan ke array, kalo skor klub B lebih besar maka nama klub B yang disimpan, dan kalo skor sama maka disimpan string "Draw". Proses input skor akan bakal jalan terus sampe salah satu atau kedua skor yang dimasukkan nilainya negatif, itu jadi tanda kalo pertandingannya udah selesai.

### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom

#### 4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0

	for {
		fmt.Scanf("%c", &ch)

		if ch == '\n' || ch == ' ' {
			continue
		}

		if ch == '.' || *n >= NMAX {
			break
		}

		t[*n] = ch
		*n++
	}
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks  : ")
	isiArray(&tab, &m)

	for i := 0; i < m; i++ {
		fmt.Printf("%c ", tab[i])
	}
	fmt.Println()

	// palindrom
	if palindrom(tab, m) {
		fmt.Println("Palindrom ? true")
	} else {
		fmt.Println("Palindrom ? false")
	}

	// reverse
	balikanArray(&tab, m)

	fmt.Print("Reverse teks : ")
	for i := 0; i < m; i++ {
		fmt.Printf("%c ", tab[i])
	}
	fmt.Println()
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul9/Output/output-soal4.png)
Penjelasan :
Program ini pakai array bertipe rune buat nyimpen karakter yang diinput user. Inputnya dibaca satu per satu, dan bakal berhenti kalau ketemu tanda titik (.).
Fungsi balikanArray pakai cara two-pointer, jadi ada dua penunjuk dari kiri dan kanan. Tiap langkah, dua elemen itu ditukar posisinya, terus yang kiri geser ke kanan dan yang kanan geser ke kiri sampai ketemu di tengah.
Nah, fungsi palindrom ngecek apakah teks itu palindrom atau bukan dengan cara bandingin karakter dari depan dan belakang secara berpasangan. Kalau ada yang beda, berarti bukan palindrom. Tapi kalau semuanya sama, berarti itu palindrom dan fungsi bakal ngembaliin nilai true.