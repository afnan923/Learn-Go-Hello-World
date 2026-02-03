package main
import "fmt"

//STRUCT
type Mahasiswa struct{
	Nama string
	Umur int
	Jurusan string
}
//terima alamat
func ubahNilai(x *int) {
	//ubah nilai di alamat itu
	*x = 10
}
func ubahTugas2(x *int){
	*x = 100
}
func ubahTugas(x ,y *int) {
	temp := *x
	*x = *y
	*y = temp
}
func ubahJurusan(m *Mahasiswa){
	m.Jurusan = "Akuntansi"
}


func main() {
	angka := 5
	angka2 := 6
	//kirim alamat angka
	ubahNilai(&angka)
	fmt.Println(angka)
	//tugas
	ubahTugas2(&angka)
	fmt.Println(angka)
	//tugas 2
	mhs := Mahasiswa{
		Nama:"Afnan",
		Umur:20,
		Jurusan:"Informatika",
	}
	ubahJurusan(&mhs)
	fmt.Println(mhs)
	//tugas 3
	ubahTugas(&angka,&angka2)
	fmt.Println(angka,angka2)
}