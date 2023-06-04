package main 

import "fmt"

func main() { 	// 643

	//var num int
	var five int
	fmt.Println("Enter the number:")
	//fmt.Scanln(&num)
	fmt.Scanln(&five)

	//fmt.Println((num % 10) + (num % 100 / 10) +  (num / 100))
	
	fmt.Println((five % 25 + 5) * (five / 50 - 5) + (525 % 35 + 10 ) / (15 % 10 )  )
}