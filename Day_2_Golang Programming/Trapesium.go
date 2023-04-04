package main

import "fmt"

func main() {
	var a, b, t float64
	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	fmt.Print("t = ")
	fmt.Scan(&t)

	Luas := (a + b) * t / 2
	fmt.Print("Luas Trapesium Adalah ", Luas)

}
