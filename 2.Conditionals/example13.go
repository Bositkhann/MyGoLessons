package main

import "fmt"

func main(){

var sekund int
  
// 			B7

fmt.Println("Necha sekund o'tganini yozing :")
fmt.Scanln(&sekund)
if sekund < 60 {
	fmt.Println("60 dan katta son kiriting ")
}
fmt.Println(sekund / 60, "minut o'tgan")

}