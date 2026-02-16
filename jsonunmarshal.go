package main
import (
	"encoding/json"
	"fmt"
)

type Mahasiswa struct {
	Nama string
	Umur int
}

func main() {
	jsonData := `{"Nama":"Afnan","Umur":20}`

	var mhs Mahasiswa
	json.Unmarshal([]byte(jsonData), &mhs)

	fmt.Println(mhs.Nama)
	fmt.Println(mhs.Umur)
}