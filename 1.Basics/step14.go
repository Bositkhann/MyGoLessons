package main

import "fmt"

func main(){ 		// A4

	var eni int
	var boyi int
	var uzunligi int

	fmt.Print("Eni:")
	fmt.Scanln(&eni)
	fmt.Print("Boyi:")
	fmt.Scanln(&boyi)
	fmt.Print("Uzunligi:")
	fmt.Scanln(uzunligi)

	fmt.Println("Javob:", eni * boyi * uzunligi)

}