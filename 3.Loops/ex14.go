package main

import "fmt"

func main(){

	var a int 

	fmt.Println("5 xonali son kiriting :")
	fmt.Scanln(&a)
fmt.Println("Sonlarning teskarisi :")
	fmt.Print(a % 10)
	fmt.Print(a / 10 % 10 )
	fmt.Print(a / 100 % 10)
	fmt.Print(a / 1000 % 10)
	fmt.Println(a / 10000 % 10)
	
	fmt.Println("Sonlarning yig'indisi :")
	fmt.Println((a % 10) + (a / 10 % 10 ) + (a / 100 % 10) + (a / 1000 % 10) + (a / 10000 % 10))

}