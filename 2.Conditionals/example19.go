package main

import "fmt"

func main(){

	// B1
	var a int
	var b int

	fmt.Println("Uchburchak katetlarini kiriting:")
	fmt.Print("a = ")
	fmt.Scanln(&a)
	fmt.Print("b = ")
	fmt.Scanln(&b)

	fmt.Print("gipotenuza =")
	fmt.Println((a * a) + (b * b))



}