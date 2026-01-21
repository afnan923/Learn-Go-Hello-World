package main

import "fmt"

func main(){
	fmt.Println("Hello,World!")

	//deklarasi variabel dengan tipe data
	var name = "Afnan"
	var umur = 12
	//deklarasi dengan tipe inferensi atau untuk koma
	version := 1.23
	//deklarasi konstanta
	const pi = 3.14
	
	//panggil semua variabel atau deklarasi
	fmt.Println("Nama",name)
	fmt.Println("Umur",umur)
	fmt.Println("Versi",version)
	fmt.Println("Nilai Pi",pi)
}