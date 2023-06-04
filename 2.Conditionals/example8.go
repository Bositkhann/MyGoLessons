package main

import "fmt"

func main (){

var a int
var b int 

fmt.Println("Enter the first number:")
fmt.Scanln(&a)
fmt.Println("Enter the second number:")
fmt.Scanln(&b)

fmt.Println("Answer:",a * b)

}