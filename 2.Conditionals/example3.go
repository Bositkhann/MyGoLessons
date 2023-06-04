package main 

import "fmt"

func main (){

	var a int 
	var b int 

	fmt.Println("Enter the number:")
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	if a > b {
		fmt.Println("a is greater than b")

	}else {
		fmt.Println("b is greater than a")
	}
	




}