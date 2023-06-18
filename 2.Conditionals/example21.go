package main

import "fmt"

func main(){

	// 			B5
	
	var a int
	var b int
	var c int

	fmt.Println("ixtiyoriy uchta son kiriting ")

	fmt.Print("1 son = ")
	fmt.Scanln(&a)
	fmt.Print("2 son = ")
	fmt.Scanln(&b)
	fmt.Print("3 son = ")
	fmt.Scanln(&c)


	if a > b && a > c  {
		fmt.Println("katta son : ",a)
	}else if b > a && b > c  {
		fmt.Println("katta son : ",b)
	}else if c > a && c > b  {
		fmt.Println("katta son : ",c)
	}else if a == b && a == c {
		fmt.Println("bu sonlar teng")
	}else {
		fmt.Println("Kalla qo'ydin")
	}

}