package main

import "fmt"

//fungsi sederhana
func greet(name string) string {
	return "Halo, " + name
}

func main() {
	message := greet("BuildWithAfnan")
	fmt.Println(message)
}