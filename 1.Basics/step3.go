package main

import "fmt"
import "time"


func main(){

	var name string //Declare --> "" bo'sh
	fmt.Scanln(&name)
	fmt.Println("Helllo -->" , name)  // Hello --> Alixan


	vaqt := time.Date(2022, time.February, 11,23,47,78,45, time.Local)
	fmt.Println(vaqt)

	fmt.Println(time.Now())	// Hozirgi vaqt 
}