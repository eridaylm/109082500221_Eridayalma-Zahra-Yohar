# <h1 align="center">Laporan Praktikum Modul 2 - Algoritma Pemrograman 2 </h1>
<p align="center">Eridayalma Zahra Yohar - 109082500221</p>

## Modul 2 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)
#### 1.go

```go
package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {

	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	p1 := permutasi(a, c)
	k1 := kombinasi(a, c)

	p2 := permutasi(b, d)
	k2 := kombinasi(b, d)

	fmt.Println(p1, k1)
	fmt.Println(p2, k2)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul3/Output/output-soal1.png)
Penjelasan : 
Program ini buat ngitung permutasi dan kombinasi dari angka yang dimasukin sama user.
Di awal ada function faktorial(n) yang gunanya buat ngitung faktorial dari suatu angka. Faktorial itu perkalian berurutan dari 1 sampai angka itu, misal 5! = 5 × 4 × 3 × 2 × 1.
Di dalem function itu, program bikin variabel hasil yang awalnya diisi 1. Terus program pakai perulangan for dari 1 sampai n. Setiap perulangan, nilai hasil dikali dengan angka i, jadi lama-lama hasilnya bakal jadi faktorial.
Terus ada function permutasi(n, r), Function ini dipake buat ngitung permutasi, yaitu banyaknya cara menyusun beberapa objek dengan memperhatikan urutan, rumusnya itu : n! / (n-r)!
Makanya di program ini manggil function faktorial buat ngitung.
Terus ada juga function kombinasi(n, r) yang dipake buat ngitung kombinasi, yaitu jumlah cara memilih objek tanpa memperhatikan urutan, rumusnya: n! / (r! × (n-r)!)

Implementasi : 
- Di bagian main, program membaca 4 buah input angka dari user yaitu a, b, c, dan d.
- Habis itu program menghitung:
    a. permutasi dan kombinasi dari a dan c
    b. permutasi dan kombinasi dari b dan d
- Terus hasilnya disimpen di variabel p1, k1, p2, dan k2, habis itu di munculin outputnya

Jadi, intinya program ini cuma minta input angka terus ngitung permutasi dan kombinasi dari angka itu.


### 2. Diberikan tiga buah fungsi matematika yaitu 𝑓 (𝑥) = 𝑥2 , 𝑔 (𝑥) = 𝑥 − 2 dan ℎ (𝑥) = 𝑥 + 1. Fungsi komposisi (𝑓𝑜𝑔𝑜ℎ)(𝑥) artinya adalah 𝑓(𝑔(ℎ(𝑥))). Tuliskan 𝑓(𝑥), 𝑔(𝑥) dan ℎ(𝑥) dalam bentuk function.
#### 2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul2/Output/output-soal2.png)
Penjelasan : 
Program ini itu tentang fungsi komposisi matematika, di soal dikasih tiga fungsi :
    a. f(x) = x²
    b. g(x) = x − 2
    c. h(x) = x + 1

Di program, tiga fungsi itu dibuat dalam bentuk function
    1. Function f(x) mengembalikan nilai x * x, jadi artinya x dipangkatkan 2.
    2. Function g(x) mengembalikan nilai x - 2.
    3. Function h(x) mengembalikan nilai x + 1.

Implementasi :
- Di bagian main, program membaca tiga angka dari input yaitu a, b, dan c.
- Habis itu program ngitung beberapa komposisi fungsi:
    a. f(g(h(a)))
    Artinya nilai a masuk dulu ke fungsi h, hasilnya masuk ke g, lalu hasilnya lagi masuk ke f.

    b. g(h(f(b)))
    Artinya nilai b masuk dulu ke f, lalu hasilnya ke h, lalu ke g.

    c. h(f(g(c)))
    Artinya nilai c masuk dulu ke g, lalu ke f, lalu ke h.

Hasil dari masing-masing perhitungan itu disimpan di variabel hasil1, hasil2, dan hasil3, habis itu dimunculin outputnya 

Jadi intinya, program ini ngejalanin fungsi matematika yang saling dimasukin satu sama lain sesuai konsep komposisi fungsi.

### 3. [Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥, 𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥, 𝑦) berdasarkan dua lingkaran tersebut.
#### 3.go

```go
package main

import "fmt"

package main

import (
	"fmt"
	"math"
)

func didalam(cx, cy, r, x, y float64) bool {
	jarak := math.Sqrt((x-cx)*(x-cx) + (y-cy)*(y-cy))
	return jarak <= r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dalam1 := didalam(cx1, cy1, r1, x, y)
	dalam2 := didalam(cx2, cy2, r2, x, y)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul2/Output/output-soal3.png)
Penjelasan :
Program ini tentang posisi sebuah titik terhadap dua lingkaran, sebuah lingkaran ditentukan oleh:
    - titik pusat (cx, cy)
    - jari-jari r

Di program dibuat function didalam() yang gunanya buat mengecek apakah sebuah titik ada di dalam lingkaran atau engga.

Cara kerjanya itu ngitung jarak antara titik (x, y) dengan pusat lingkaran (cx, cy) pake rumus jarak: 
- √((x - cx)² + (y - cy)²), rumus ini dihitung pake math.Sqrt.

Kalau jarak titik ke pusat lingkaran lebih kecil atau sama dengan radius, berarti titik itu berada di dalam lingkaran. Makanya function ini mengembalikan nilai true atau false.

Implementasi : 
- Di bagian main, program ngebaca input dari user yaitu:
    1. pusat dan radius lingkaran pertama
    2. pusat dan radius lingkaran kedua
    3. koordinat titik (x, y)

- Setelah itu program ngecek apakah titik tersebut berada di dalam lingkaran pertama dan kedua dengan memanggil function didalam().
- Hasilnya disimpan di variabel dalam1 dan dalam2.
- Terus program pake if else buat nentuin posisi titik:
    a. kalau dalam1 dan dalam2 true → titik ada di dalam kedua lingkaran
    b. kalau cuma dalam1 true → titik ada di lingkaran 1
    c. kalau cuma dalam2 true → titik ada di lingkaran 2
    d. kalau dua-duanya false → titik ada di luar kedua lingkaran
    e. Terakhir program munculin outputnya

Jadi intinya program ini ngecek apakah suatu titik berada di dalam lingkaran 1, lingkaran 2, keduanya, atau di luar semuanya.