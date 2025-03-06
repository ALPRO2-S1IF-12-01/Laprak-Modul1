package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung luas permukaan tabung
func luasPermukaanTabung(r, t float64) float64 {
	return 2 * math.Pi * r * (r + t)
}

// Fungsi untuk menghitung volume tabung
func volumeTabung(r, t float64) float64 {
	return math.Pi * math.Pow(r, 2) * t
}

func main() {
	var r, t float64

	// Input jari-jari dan tinggi tabung dengan validasi
	fmt.Print("Masukkan jari-jari tabung: ")
	_, errR := fmt.Scan(&r)
	fmt.Print("Masukkan tinggi tabung: ")
	_, errT := fmt.Scan(&t)

	// Memeriksa apakah input valid
	if errR != nil || errT != nil {
		fmt.Println("Input tidak valid! Harap masukkan angka yang benar.")
		return
	}

	// Memeriksa apakah jari-jari dan tinggi bernilai positif
	if r <= 0 || t <= 0 {
		fmt.Println("Jari-jari dan tinggi tabung harus lebih dari nol.")
		return
	}

	// Menghitung luas permukaan dan volume
	luas := luasPermukaanTabung(r, t)
	volume := volumeTabung(r, t)

	// Menampilkan hasil
	fmt.Println("===================================")
	fmt.Printf("Luas Permukaan Tabung: %.2f satuan²\n", luas)
	fmt.Printf("Volume Tabung: %.2f satuan³\n", volume)
	fmt.Println("===================================")
}
