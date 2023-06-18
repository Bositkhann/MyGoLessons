package main

import "fmt"

func main(){

var sekund int
  


fmt.Println("Necha sekund o'tganini yozing :")
fmt.Scanln(&sekund)
if sekund < 3600 {
	fmt.Println("3600 dan katta son kiriting ")
}
fmt.Println(sekund / 3600, "soat o'tgan")

}