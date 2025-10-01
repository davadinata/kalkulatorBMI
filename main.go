package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// serve ke folder static
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// endpoint / bmi
	http.HandleFunc("/bmi", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		// ambil data
		weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
		height, _ := strconv.ParseFloat(r.FormValue("height"), 64)

		// cek val
		if height <= 0 {
			http.Error(w, "Tinggi badan harus lebih dari 0", http.StatusBadRequest)
			return
		}
		if weight <= 0 {
			http.Error(w, "Berat badan harus lebih dari 0", http.StatusBadRequest)
			return
		}

		// calc BMI
		imt := weight / ((height / 100) * (height / 100))

		result := fmt.Sprintf("Indeks Massa Tubuh (IMT): %.2f\n", imt)

		// val kategori
		if imt < 18.5 {
			result += "Kategori: Berat badan kurang"
		} else if imt >= 18.5 && imt < 24.9 {
			result += "Kategori: Ideal"
		} else if imt >= 25.0 && imt < 29.9 {
			result += "Kategori: Berat badan berlebih"
		} else {
			result += "Kategori: Obesitas"
		}

		fmt.Fprint(w, result)
	})

	// start server
	fmt.Println("Server jalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}