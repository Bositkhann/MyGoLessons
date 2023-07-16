package main

import "fmt"

func main() {
//                      B13
 var a int
 var b int
 var c int

 fmt.Scanln(&a) 
 fmt.Scanln(&b) 
 fmt.Scanln(&c) 


 if a+b > c && a+c > b && b+c > a { 
  fmt.Println("Ha")
 } else {
  fmt.Println("yo'q")
 }


}