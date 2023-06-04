package main

import "fmt"

func main(){

	/*
		Evaluation system
		80 --> 100 --> 5
		60 --> 80 --> 4
		40 --> 60 --> 3
		Below from 40 is failed
	*/

	var one int 
	fmt.Println("Enter your ball:")
	fmt.Scanln(&one)
	if one > 100 {
		fmt.Println("Oh, this ball is very  good 👍")
	} else if 80 <= one && one <= 100 {
		fmt.Println("Your ball 5")
	} else if 60 <= one && one <= 80 {
		fmt.Println("Your ball 4")
	} else if 40 <= one && one <= 60 {
		fmt.Println("Your ball 3")
	} else {
		fmt.Println("Your ball 0 🤦‍♂️")
	}
}
