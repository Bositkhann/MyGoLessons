package main

import "fmt"

func main(){

	i := 1000
	for  i >= 500 {
		fmt.Println("==>", i)
		i -= 1
	}
}