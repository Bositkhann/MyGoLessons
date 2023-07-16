package main

import "fmt"

func main(){

var sekund int
//					B7 . 3

fmt.Println("Necha sekund vaqt o'tganini yozing:")
fmt.Scanln(&sekund)

fmt.Print(sekund / 60, " minut va ")
fmt.Println(sekund % 60 ,"sekund vaqt o'tgan")




}