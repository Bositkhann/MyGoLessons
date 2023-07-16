package main

import "fmt"

func main(){


	//                    B11
	
var a int

fmt.Println("1 dan 7 gacha son kiriting :")
fmt.Scanln(&a)

if a == 1 {
	fmt.Println("1 - Dushanba ")
}else if a == 2 {
	fmt.Println("2 - Seshanba ")
}else if a == 3 {
	fmt.Println("3 - Chorshanba ")
}else if a == 4 {
	fmt.Println("4 - Payshanba ")
}else if a == 5 {
	fmt.Println("5 - Juma ")
}else if a == 6 {
	fmt.Println("6 - Shanba ")
}else if a == 7 {
	fmt.Println("7 - Yakshanba ")
}else {
	fmt.Println("Bunday kun yo'q")
}



}