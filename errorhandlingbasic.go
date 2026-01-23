package main

import (
	"errors"
	"fmt"
)

//fungsi yang mengembalikan error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Pembagian dengan nol tidak diperbolehkan")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}else {
		fmt.Println("Hasil:",result)
	}
}