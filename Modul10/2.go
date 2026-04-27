package main

import "fmt"

func main() {
	var jmlikan int
	var ikanperwadah int
	var beratIkan [1000]float64

	fmt.Print("Masukan jumlah ikan (x) dan ikan per wadah (y): ")
	fmt.Scan(&jmlikan, &ikanperwadah)

	for i := 0; i < jmlikan; i++ {
		fmt.Printf("Masukan berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	jmlwadah := jmlikan / ikanperwadah
	if jmlikan%ikanperwadah != 0 {
		jmlwadah++
	}

	beratperwadah := make([]float64, jmlwadah)

	for i := 0; i < jmlikan; i++ {
		nomorWadah := i / ikanperwadah
		beratperwadah[nomorWadah] += beratIkan[i]
	}

	fmt.Print("Total berat per wadah: ")
	ttlsmuawadah := 0.0
	for i, total := range beratperwadah {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", total)
		ttlsmuawadah += total
	}
	fmt.Println()

	rataRataWadah := ttlsmuawadah / float64(jmlwadah)
	fmt.Printf("Rata-rata berat per wadah: %.2f\n", rataRataWadah)
}