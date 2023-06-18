package main

import "fmt"

func main(){

	// 			B9
	
	var a int
	var b int
	var c int
	var d int
	var e int

	fmt.Println("ixtiyoriy beshta son kiriting ")
	fmt.Print("1 son = ")
	fmt.Scanln(&a)
	fmt.Print("2 son = ")
	fmt.Scanln(&b)
	fmt.Print("3 son = ")
	fmt.Scanln(&c)
	fmt.Print("4 son = ")
	fmt.Scanln(&d)
	fmt.Print("5 son = ")
	fmt.Scanln(&e)

	if a > b && a > c && a > d && a > e {
		fmt.Println("katta son : ",a)
	}else if b > a && b > c && b > d && b > e  {
		fmt.Println("katta son : ",b)
	}else if c > a && c > b && c > d && c > e {
		fmt.Println("katta son : ",c)
	}else if d > a && d > c && d > b && d > e  {
		fmt.Println("katta son : ",d)
    }else if e > a && e > b && e > d && e > c {
		fmt.Println("katta son : ",e)
	}else { 
		fmt.Println("bu sonlar teng")
	}

}