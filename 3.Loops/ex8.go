package main

import "fmt"

func main(){

	var i int = 0 

	for {

		if 7 == i {
			break
		}
		if i <= 10 {
			fmt.Println("loop", i)

		}
		i++
	}
}