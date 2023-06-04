package main

import "fmt"

func main(){

	a := 18
	b := 10

	if a + b < 20 {
		fmt.Println("the number 20 is greater than a + b")
	}
	if a + b > 20 {
		fmt.Println("the number a + b is greater than 20")
	}

}