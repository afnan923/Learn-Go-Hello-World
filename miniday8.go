package main
import "fmt"
import "sort"

func main() {
	//slice 
	a := []int{1,2,3}
	b := make([]int, len(a))
	copy(b, a)
	b[0] = 100

	fmt.Println(a)
	//slice int
	angka := []int{5,2,9,1}
	sort.Ints(angka)
	fmt.Println(angka)
	//slice string
	nama := []string{"Citra","Afnan","Budi"}
	sort.Strings(nama)
	fmt.Println(nama)

	//map cek key ada atau tidak
	nilai := map[string]int{
		"Afnan": 90,
	}

	skor, ada := nilai["Afnan"]

	if ada {
		fmt.Println("Ada:", skor)
	} else {
		fmt.Println("Tidak Ada")
	}

	//nested map
	data := map[string]map[string]int{
		"Afnan": {
			"Math": 90,
			"Go": 95,
		},
	}
	fmt.Println(data["Afnan"]["Go"])

	//slice of struct
	// type Mahasiswa struct{
	// 	Nama string
	// }

	// list := []Mahasiswa{
	// 	{"Afnan"},
	// 	{"Dani"},
	// }
	// //loop
	// for _, m := range list {
	// 	fmt.Println(m.Nama)
	// }

	//TUGAS 1
	angkatugas := []int{91,14,71,22,1}
	sort.Ints(angkatugas)
	fmt.Println(angkatugas)

	//TUGAS2
	nilaitugas := map[string]int{
		"Afnan": 80,
	}
	skor, adaTugas := nilaitugas["Afnan"]

	if adaTugas{
		fmt.Println("Ada:", skor)
	} else {
		fmt.Println("Tidak ada")
	}

	//TUGAS3
	type MahasiswaBaru struct{
		Nama string
	}

	list := []MahasiswaBaru{
		{"Afnan"},
		{"Dani"},
		{"Citra"},
	}
	//loop
	for _, m := range list {
		fmt.Println(m.Nama)
	}
}