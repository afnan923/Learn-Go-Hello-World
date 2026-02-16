package main
import (
	"encoding/json"
	"fmt"
)

type Mahasiswa struct {
	Nama string `json:"nama"`
	Umur int `json:"umur"`
	Jurusan string `json:"jurusan"`
}

func main() {
	pilih := "unmarshal"
	if pilih == "marshal" {
		mhs := Mahasiswa{"Afnan", 20, "Informatika"}

		data, _ := json.Marshal(mhs)
		fmt.Println(string(data))
	} else if pilih == "unmarshal" {
		jsonData := `{"nama":"Afnan","umur":20,"jurusan":"Informatika"}`

		var mhs Mahasiswa
		json.Unmarshal([]byte(jsonData), &mhs)

		fmt.Println(mhs.Nama)
		fmt.Println(mhs.Umur)
		fmt.Println(mhs.Jurusan)
	} else {
		mhs := Mahasiswa{"Afnan", 20, "Informatika"}

		data, _ := json.Marshal(mhs)
		fmt.Println(string(data))
	}
	
}