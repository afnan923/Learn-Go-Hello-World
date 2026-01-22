package main

import "fmt"

func main(){
	//if statement
	num :=10
	if num > 5 {
		fmt.Println("Angka lebih besar dari 5")
	} else {
		fmt.Println("Angka kurang atau sama dengan 5")
	}

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
}