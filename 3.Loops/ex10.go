package main

import ("fmt"
		"time"
)
func main(){

	var number int = 0
	for number <= 11{
		fmt.Println("===>", number)
		number ++
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Next time ")

	fmt.Printf("\n")
}