package main
import(
	"encoding/json"
	"fmt"
)

type Mahasiswa struct{
	Nama string 
	Umur int
}

func main() {
	mhs := Mahasiswa{"Afnan", 20}

	data, _ := json.Marshal(mhs)
	fmt.Println(string(data))
}