package main

import "fmt"

func main() {

	

	var num int
	fmt.Print("Enter the number:")
	fmt.Scanln(&num)

	
	fmt.Println("===>", num / 100 % 10)
	fmt.Println("===>", num / 10 % 10 )
	fmt.Println("===>", num % 10)

}