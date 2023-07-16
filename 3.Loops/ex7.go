package main

import "fmt"

func main(){

	for a := 0; a <= 100 ; a++ {
		if a > 55 {
			break
		}
		fmt.Println(a)
	}
}