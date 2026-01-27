package main

import "fmt"

func main(){
	fmt.Println("Hello,World!")

	//deklarasi variabel dengan tipe data
	// var name = "Afnan" = ini untuk digunakan tanpa tipe data
	var nama string = "Dani"
	// var umur = 12
	var umurku int = 21 // ini digunakan dengan tipe data
	//short declaration
	mobil := 2.0
	harisenin := true
	var hobi = "Hobi saya adalah mengisi waktu luang dengan apapun itu"
	var motivasi = "Motivasi saya adalah ketika seseorang sedang membicarakan dan merendahkan saya"
	//deklarasi dengan tipe inferensi atau untuk koma
	version := 1.23
	//deklarasi konstanta
	const pi = 3.14
	//tipe data int untuk di database dan default
	var a int = 10 // default
	var b int64 = 100 //untuk data besar(database,ID)
	//float
	var harga float64 = 99.99
	//go tidak otomatis konversi
	var x int = 100
	var y float64 = 9.9
	hasil := float64(x) + y
	//tipe boolean
	var isLogin bool = true
	//panggil semua variabel atau deklarasi
	fmt.Println("Nama=",nama)
	fmt.Println("Umur=",umurku)
	fmt.Println(mobil)
	fmt.Println(harisenin)
	fmt.Println("hobi=",hobi)
	fmt.Println("motivasi=",motivasi)
	fmt.Println("Versi=",version)
	fmt.Println("Nilai Pi=",pi)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(harga)
	fmt.Println(hasil)
	fmt.Println("Login :",isLogin)
}