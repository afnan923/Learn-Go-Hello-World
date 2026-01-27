package main

import "fmt"

func main() {
	//for umum
	for i:= 1; i <= 5; i ++{
		fmt.Println(i)
	} 
	//for sebagai while
	k := 1
	for k<= 5 {
		fmt.Println(k)
		k++
	}
	// //for tanpa kondisi atau disebut infinite loop
	// for {
	// 	fmt.Println("Loop Terus")
	// }
	//break untuk menghentikan loop
	for j := 1; j <= 10; j++ {
		if j == 5 {
			break
		}
		fmt.Println(j)
	}
	//continue untuk lewati satu iterasi
	for l := 1; l <= 5; l++ {
		if l == 3 {
			continue
		}
		fmt.Println(l)
	}
	//cetak bilangan genap
	for o := 1; o <= 10; o++ {
		if o%2 !=0 {
			continue
		}
		fmt.Println(o)
	}
	//hitung total nilai
	total := 0

	for t := 1; t <= 5; t++{
		total += t
	}
	fmt.Println("Total:", total)

	//tugas
	for g := 1; g <= 20; g++ {
		if g == 15 {
			break
		}
		if g%2 !=0 {
			continue
		} 
		fmt.Println(g)
	}
}
