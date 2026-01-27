package main 
import "fmt"

func main() {
	//array ukuran tetap dan tidak bisa nambah data
	var angka [3]int = [3]int{10, 20, 30}
	//jika kita print akan memunculkan semua angka
	fmt.Println(angka)
	//ini akan print salah satu angka dari yang diatas
	fmt.Println(angka[0])//angka 10 karena hitungannya dari 0 baru ke yang lain

	//Slice untuk menambah barang kapn saja
	buah := []string{"Apel", "Mangga"}
	//ini untuk tambah data ke slice
	buah = append(buah,"Jeruk")
	fmt.Println(buah)
	//untuk mengetahui sebuah panjang data menggunakan len
	fmt.Println(len(buah))
	//mengambil data dari slice
	fmt.Println(buah[0])//Apel

	//Loop Slice
	//ini biasanya for biasa cara manual
	for i := 0; i < len(buah); i++ {
		fmt.Println(buah[i])
	}
	//ini for range untuk mengetahui tipe list daftar
	for i, v := range buah {
		fmt.Println(i, v)
	}
	//ini untuk tidak menampilkan angka dari sebelumnya. _ = artinya tidak peduli dengan variabel ini
	for _, v := range buah {
		fmt.Println(v)
	}

	//Slice of int
	nilai := [] int{80, 75, 90}
	nilai = append(nilai,85, 70)
	total := 0
	for _, n := range nilai{
		total += n
	}

	rataRata := total / len(nilai)
	fmt.Println("Nilai:", nilai)
	fmt.Println("Total:", total)
	fmt.Println("Rata-Rata:", rataRata)
}