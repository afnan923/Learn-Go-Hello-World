package main 
import(
	"fmt" 
	"os"
)

func main() {
	mode := "read"
	if mode == "write" {
		file, err := os.Create("nama.txt")

		if err != nil {
		fmt.Println("Error:", err)
		return
		}

		defer file.Close()

		file.WriteString("Nama: Afnan\n")
		file.WriteString("Belajar Go")
	
	} else if mode == "append" {
		file, err := os.OpenFile("nama.txt", os.O_APPEND|os.O_WRONLY,0644)

		if err != nil {
		fmt.Println("Error:", err)
		return
		}
		
		defer file.Close()

		file.WriteString("\nDay 13 Selesai")
	} else {
		nama, err := os.ReadFile("nama.txt")

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(string(nama))
	}

}