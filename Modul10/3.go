package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(dataBerat arrBalita, jumlahData int, beratMin, beratMax *float64) {
	*beratMin = dataBerat[0]
	*beratMax = dataBerat[0]

	for i := 1; i < jumlahData; i++ {
		if dataBerat[i] < *beratMin {
			*beratMin = dataBerat[i]
		}
		if dataBerat[i] > *beratMax {
			*beratMax = dataBerat[i]
		}
	}
}

func rerata(dataBerat arrBalita, jumlahData int) float64 {
	totalBerat := 0.0
	for i := 0; i < jumlahData; i++ {
		totalBerat += dataBerat[i]
	}
	return totalBerat / float64(jumlahData)
}

func main() {
	var databalita arrBalita
	var jmlbalita int

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&jmlbalita)

	for i := 0; i < jmlbalita; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&databalita[i])
	}

	var beratMin, beratMax float64
	hitungMinMax(databalita, jmlbalita, &beratMin, &beratMax)

	rataRata := rerata(databalita, jmlbalita)

	fmt.Printf("Berat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
}