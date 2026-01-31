package main
import "fmt"

//STRUCT
type Mahasiswa struct{
	Nama string
	Umur int
	Jurusan string
}

func main() {
	//MAP
	nilai := map [string]int{
		"Afnan": 90,
		"Dani": 85,
		
		
	}
	//untuk menambah data dari sebuah map
	nilai["Citra"] = 88
	//untuk menghapus data
	delete(nilai, "Dani")
	//output atau akses
	fmt.Println(nilai)

	//loop dan ini berfungsi untuk mencetak satu per data dan mengambil semua data yang ada diatas
	for nama, skor := range nilai{
		fmt.Println(nama, skor)
	}


	mhs := Mahasiswa{
		Nama:"Afnan",
		Umur:20,
		Jurusan:"Informatika",
	}
	fmt.Println(mhs.Nama)
	fmt.Println(mhs.Umur)
	fmt.Println(mhs.Jurusan)

	//GABUNGAN dari MAP dan STRUCT
	mahasiswa := map [string]Mahasiswa{
		"001": {Nama: "Afnan", Umur : 20, Jurusan: "Informatika"},
		"002": {Nama: "Citra", Umur : 19, Jurusan: "Akuntansi"},
	}
	fmt.Println(mahasiswa["001"]) 

	for id, data := range mahasiswa {
	fmt.Println("ID:", id)
	fmt.Println("Nama:", data.Nama)
	fmt.Println("Umur:", data.Umur)
	fmt.Println("Jurusan:", data.Jurusan)
	fmt.Println() // spasi antar data
}

}