package main

import "fmt"

func main(){

	for a := 0; a <= 100 ; a++ {
		if a %  3== 0 {
			continue
		}
		fmt.Println(a)
	}
}