package main

import "fmt"

func main(){

	//				B4
	
	var a int
	var b int

	fmt.Println("ixtiyoriy ikkita son kiriting ")
	fmt.Print("1 son = ")
	fmt.Scanln(&a)
	fmt.Print("2 son = ")
	fmt.Scanln(&b)

	if a > b {
		fmt.Println("katta son : ",a)
	}else if a < b {
		fmt.Println("katta son : ",b)
	}else { 
		fmt.Println("bu sonlar teng")
	}




}