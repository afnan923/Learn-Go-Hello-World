package main 
import "fmt"


//Function Tanpa Return
func sapa(){
	fmt.Println("Hallo Afnan")
}

//Function Dengan Parameter
func sapaNama(nama string){
	fmt.Println("Halo", nama)
}

//Function Dengan Return
func tambah(a int, b int) int {
	return a+b
}

//multiple return
func hitung(c, d int) (int, int) {
	return c + d, c - d
}

//return error
func bagi(e, f float64) (float64, error) {
	if f == 0 {
		return 0, fmt.Errorf("tidak bisa bagi nol")
	}
	return e / f, nil
}

func kali(a int, b int) int {
	return a*b
}

func cekUmur(umur int)string{
	if umur <= 13 {
		return "Anak"
	} else if umur <= 17 {
		return "Remaja"
	} else {
		return "Dewasa"
	}
}

func luasPersegi(sisi int) int {
	return sisi * sisi
}


func main() {
	//panggil func tanpa return
	sapa()
	//panggil func dengan parameter
	sapaNama("Afnan")
	sapaNama("Dani")
	//pemanggilan func dengan return
	hasil := tambah(5,3)
	fmt.Println(hasil)
	//Pemanggilan multiple return
	tambah, kurang := hitung(10, 5)
	fmt.Println(tambah)
	fmt.Println(kurang)
	//pemanggilan return error
	hasilbagi, err := bagi(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasil:", hasilbagi)
	}
	//pemanggilan kali
	hasilkali := kali(1,3)
	fmt.Println(hasilkali) 
	//pemanggilan umur
	hasilumur := cekUmur(12)
	fmt.Println(hasilumur)
	//pemanggilan luas
	hasilluas := luasPersegi(1)
	fmt.Println(hasilluas)
}