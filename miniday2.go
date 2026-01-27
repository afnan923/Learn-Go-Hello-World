package main

import "fmt"

func main() {
	nama := "Afnan"
	tahunlahir := 2005
	tahunsekarang := 2026
	umursekarang := tahunsekarang - tahunlahir
	mahasiswa := true
	umur := 21
	tinggi := 165.5
	menikah := false

	fmt.Println("Nama:",nama)
	fmt.Println("Tahun Lahir:",tahunlahir)
	fmt.Println("Tahun Sekarang:",tahunsekarang)
	fmt.Println("Umur Sekarang:",umursekarang)
	fmt.Println("Mahasiswa:",mahasiswa)
	fmt.Println("Umur:",umur)
	fmt.Println("Tinggi Badan:",tinggi)
	fmt.Println("Menikah:",menikah)

	umur5tahun := umur + 5
	fmt.Println("Umur saya 5 Tahun lagi:",umur5tahun)
}