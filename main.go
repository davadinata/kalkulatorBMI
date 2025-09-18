package main

import "fmt"

func main() {
    var bb, tb float64

    fmt.Print("Masukan berat badan (kg): ")
    fmt.Scan(&bb)
	// fmt.Println("Debug bb:", bb)


    fmt.Print("Masukan tinggi badan (m): ")
    fmt.Scan(&tb)
	// fmt.Println("Debug tb:", tb)

    if tb <= 0 {
        fmt.Println("Tinggi badan harus lebih dari 0")
        return
    }

	if bb <= 0 {
		fmt.Println("Berat badan harus lebih dari 0")
		return
	}

    imt := bb / (tb * tb)
    fmt.Printf("Indeks Massa Tubuh (IMT) Anda: %.2f\n", imt)

	if imt < 18.5 {
		fmt.Println("Berat badan kurang")
	} else if imt >= 18.5 && imt < 24.9 {
		fmt.Println("Ideal") 
	} else if imt >= 25.0 && imt < 29.9 {
		fmt.Println("Berat badan berlebih") }
	}
