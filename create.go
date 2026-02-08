package main
import (
	"fmt"
	"os"
)

func main() {
	//untuk membuat file
	// file, err := os.Create("data.txt")

	//untuk membaca file
	data, err := os.ReadFile("data.txt")

	//untuk tambah isi file tanpa hapus yang lama
	// file, _ := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	// //tutup otomatis
	// defer file.Close()

	// //tulis teksnya
	// file.WriteString("Halo Afnan\n")
	// file.WriteString("Belajar Golang File I/O")

	//untuk membaca sebuah data
	fmt.Println(string(data))
	
	//untuk membuat sebuah data baru atau isi file baru
	// file.WriteString("\nBaris Baru")
}