package main

import "fmt"

func main() {

	//			B8
	
	var num int

	fmt.Print("Uch xonali son kiriting:")
	fmt.Scanln(&num)

	
	fmt.Println( (num / 100 % 10) + ( num / 10 % 10) + (num % 10))

}