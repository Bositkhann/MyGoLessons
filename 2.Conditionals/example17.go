package main

import "fmt"

func main(){

var sekund int
//								B7 . 4

fmt.Println("Necha sekund vaqt o'tganini yozing:")
fmt.Scanln(&sekund)

fmt.Println(sekund / 3600, " soat va ")
fmt.Println(sekund % 3600 ,"sekund vaqt o'tgan")
}