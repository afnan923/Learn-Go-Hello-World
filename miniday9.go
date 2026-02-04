package main

import "fmt"

//struct itu cuma data
type Mahasiswa struct{
	Nama string
}

//interface
type Hewan interface {
	suara()
}
type Kucing struct{}

//interface tugas 3
type BangunDatar interface{
	luas()
}
type Persegi struct{}
type Lingkaran struct{}

//tugas 1
type Mobil struct{
	Merk string
}
//tugas 1 bagian method
func (mm Mobil) jalan() {
	fmt.Println(mm.Merk,"berjalan")
}
//tugas 2 bagian pointer method
func (mm *Mobil) gantiMerk() {
	mm.Merk = "Mobil Pajero"
	fmt.Println(mm.Merk, "berjalan")
}

//method dasar
func (m Mahasiswa) sapa() {
	fmt.Println("Halo", m.Nama)
}
//method ubah data menggunakan pointer method
func (m *Mahasiswa) gantiNama() {
	m.Nama = "Dani"
	fmt.Println("Nama:", m.Nama)
}
//panggil func kucing
func (k Kucing) suara() {
	fmt.Println("Meong")
}
//panggil tugas 3
func (p Persegi) luas(){
	fmt.Println("Persegi")
}
func (l Lingkaran) luas(){
	fmt.Println("Lingkaran")
}
func main() {
	//panggil method mahasiswa
	mhs := Mahasiswa{"Afnan"}
	mhs.sapa()
	mhs.gantiNama()
	//panggil interface
	var h Hewan
	h = Kucing{}
	h.suara()
	//panggil method tugas1
	mbl := Mobil{"Mobil Avanza"}
	mbl.jalan()
	mbl.gantiMerk()
	//panggil tugas 3
	var b BangunDatar
	b = Persegi{}
	b = Lingkaran{}
	b.luas()

}