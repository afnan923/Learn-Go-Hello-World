package main
import (
	"fmt"
	"time"
)

func cetakA() {
	for i := 1; i <= 6; i++ {
		fmt.Println( i)
		time.Sleep(1 * time.Second)
	}
}

func cetakB() {
	for h := 'A'; h <= 'F'; h++ {
		fmt.Println(string(h))
		time.Sleep(1 * time.Second)
	}
}


func main() {
	go cetakA()
	go cetakB()

	time.Sleep(7 * time.Second)
}