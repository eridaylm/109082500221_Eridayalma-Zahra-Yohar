# <h1 align="center">Laporan Praktikum Modul 2 - Algoritma Pemrograman 2 </h1>
<p align="center">Eridayalma Zahra Yohar - 109082500221</p>

## Modul 5 

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan 𝑆𝑛 = 𝑆𝑛−1 + 𝑆𝑛−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.

#### 1.go

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan N : ")
	fmt.Scan(&n)
	fmt.Printf("Fibonacci ke-%d = %d\n", n, fibonacci(n))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul5/Output/output-soal1.png)
Penjelasan : 
Program ini ngitung suku ke-N dari deret Fibonacci pake fungsi rekursif. Deret Fibonacci didefinisikan dengan suku ke-0 = 0, suku ke-1 = 1, dan suku berikutnya adalah jumlah dua suku sebelumnya. Fungsi fibonacci(n) bekerja dengan cara: jika n=0 kembalikan 0, jika n=1 kembalikan 1, selain itu kembalikan fibonacci(n-1) + fibonacci(n-2). Cara kerjanya user masukkin nilai N, terus program manggil fungsi fibonacci(N) baru nampilin hasilnya.

### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.
#### 2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul5/Output/output-soal2.png)
Penjelasan : 
Program ini nampilin pola segitiga bintang pake rekursif. User masukkin N, program nyetak N baris dimana baris ke-i berisi i buah bintang. Program punya dua prosedur rekursif: cetakbintang(k) yang mencetak bintang sebanyak k kali dengan cara mencetak satu bintang lalu memanggil dirinya buat k-1, dan pola(barissekarang, ttlbaris) yang mencetak bintang untuk baris sekarang lalu memanggil dirinya buat baris berikutnya hingga mencapai ttlbaris. Cara kerjanya user masukkin N terus program memanggil pola(1, N).

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.

#### 3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/eridaylm/10902500221_Eridayalma/blob/main/Modul5/Output/output-soal3.png)
Penjelasan :
Program ini nampilin semua faktor bilangan N pake rekursif. Faktor adalah bilangan yang dapat membagi N habis (tanpa sisa). Prosedur tampilfaktor(calon, bilangan) bekerja dengan mengecek apakah bilangan habis dibagi calon, kalo iya nanti calon dicetak sebagai faktor. Habis itu prosedur manggil dirinya sendiri dengan calon+1, dan berhenti ketika calon melebihi bilangan. Cara kerjanya user masukkin N, program manggil tampilfaktor(1, N) terus nyetak semua faktor dari 1 sampai N secara urut.