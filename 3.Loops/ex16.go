package main

import "fmt"

func main(){

	var a int = 13
	var b int 


	for true{
		 fmt.Println("0 dan 100 gacha son tanlang")
		fmt.Scanln(&b)
		if b == a {
			fmt.Println("Siz topdingiz")
			break
		}else if b > 100 {
			fmt.Println("Xato")
			
		}else {
			fmt.Println("Siz topa olmadingiz , yana bir marta sinab ko'ring")
		}
	}
}