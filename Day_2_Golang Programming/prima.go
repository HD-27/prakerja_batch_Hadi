package main

import "fmt"

func main() {
	var angka int
	fmt.Println("Masukkan Angka")
	fmt.Scan(&angka)
	if angka <= 1 {
		fmt.Print(angka, "Bukan Prima")
	} else {
		Prima := true
		for i := 2; i <= angka/2; i++ {
			if angka%i == 0 {
				Prima = false
				break
			}
		}
		if Prima {
			fmt.Println(angka, "adalah bilangan prima")
		} else {
			fmt.Println(angka, "bukan bilangan prima")
		}
	}
}
