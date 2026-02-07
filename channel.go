package main
import (
	"fmt"
	"time"
)
func kirim(chs chan string){
	time.Sleep(2 * time.Second)
	chs <- "Halo Afnan"
}
func tugas1(ch1 chan int){
	ch1 <- 10
}
func tugas2(ch2 chan string){
	time.Sleep(2 * time.Second)
	ch2 <- "Halo Go"
}
func cetakA(ch3 chan string) {
	for i := 1; i <= 3; i++ {
		ch3 <- fmt.Sprint(i)
	}
}

func cetakB(ch3 chan string) {
	for h := 'A'; h <= 'C'; h++ {
		ch3 <- string(h)
	}
}

func main() {
	//bentuk dasar channel
	ch := make(chan int)
	chs := make(chan string)
	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan string)

	//memanggil func diatas
	go kirim(chs)
	go tugas1(ch1)
	go tugas2(ch2)
	go cetakA(ch3)
	go cetakB(ch3)

	go func() {
		//kirim data ke channel
		ch <- 5
	}()
	//terima data dari channel
	hasil := <-ch
	fmt.Println(hasil)
	//terima data dari channel string
	pesan := <-chs
	fmt.Println(pesan)
	//terima data dari channel tugas 1
	hasiltugas1 := <-ch1
	fmt.Println(hasiltugas1)
	//terima data dari channel tugas 1
	hasiltugas2 := <-ch2
	fmt.Println(hasiltugas2)
	//terima data dari channel tugas 3
	for i := 0; i <6; i++{
		fmt.Println(<-ch3)
	}
}