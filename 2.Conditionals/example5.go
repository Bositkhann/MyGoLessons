package main

import "fmt"

func main(){

	fmt.Println("1.Your balance \n 2.Your SMS \n 3.Your Limit")
	var choice int 
	fmt.Println("Choose number: ")
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		fmt.Println("Your balance is 9999999999999$")
		break	// --> break or NOT NEEDED
	case 2:
		fmt.Println("999 SMS came to you from president ")
		break
	case 3:
		fmt.Println("Limit 999 TB")
		break
	default:
		fmt.Println("This number is No! DANG 🤦‍♂️")
	}




}