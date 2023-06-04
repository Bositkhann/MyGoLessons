package main

import "fmt"

func main () {

box := "Hello"
fmt.Println(box) // bunda faqat var ni o'zi bo'ladi

//	fmt.Printf("", box)  -->  bunda 2 ta parametr oladi "" va box(var)
fmt.Printf("Qalle: %s \n", box)
fmt.Printf("Salom: %v \n", box)


fmt.Printf("%v %v %v %v \n", 11, 22.5, "Alixan", true)



x,y,z,b := 1, 77.6, "World", false
fmt.Printf("1: %T\n2: %T\n3: %T\n4: %T\n", x, y, z, b)



















}