package main

import "fmt"

func main() {
	var angka int
	fmt.Println("Masukkan Angka")
	fmt.Scan(&angka)

	if angka%7 == 0 {
		fmt.Print(angka, " Adalah Kelipatan 7")
	} else {
		fmt.Print(angka, " Bukan Kelipatan 7")
	}
}
