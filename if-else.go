package main

import "fmt"

func main () {
	//if statement
	// nilai := 74
	// if nilai >= 80 {
	// 	fmt.Println("Dapet Nilai A")
	// } else if nilai >= 75 {
	// 	fmt.Println("Dapet Nilai B")
	// } else {
	// 	fmt.Println("Dapet Nilai C")
	// }

	// if umur := 21; umur >= 20{
	// 	fmt.Println("Dewasa")
	// }

	//for loop
	for i :=1; i < 5; i++ {
		fmt.Println("Iterasi ke:", i)
	}

	//switch case
	day := "Senin"
	switch day{
	case "Senin":
		fmt.Println("Hari ini Senin")
	case "Selasa":
		fmt.Println("Hari ini Selasa")
	default:
		fmt.Println("Hari tidak diketahui")
	}

	nilai := 75
	if nilai >= 75 {
		fmt.Println("Lulus")
	}else {
		fmt.Println("Tidak Lulus")
	}

	umur := 20
	if umur <= 13 {
		fmt.Println("Anak-Anak")
	}else if umur <= 17{
		fmt.Println("Remaja")
	}else {
		fmt.Println("Dewasa")
	}

}