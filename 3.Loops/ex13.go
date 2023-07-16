package main

import "fmt"

func main(){

// 						B12

	var a int

	fmt.Println("Bahoyini kirit: ")
	fmt.Scanln(&a)
	if a > 100 {
		fmt.Println("San o'zini kim dib o'ylavossan a ?")
	}else if a <= 100 && a >= 80 {
		fmt.Println("Bahoyin 5 ")
	}else if a <= 80 && a >= 60 {
		fmt.Println("Bahoyin 4 ")
	}else if a <= 60 && a >= 40 {
		fmt.Println("Bahoyin 3 ")
	}else if a <= 40 && a >= 20 {
		fmt.Println("Bahoyin 2 ")
	}else if a < 20 {
		fmt.Println("Keyingi safar kalleni ishlatib javob bersen baho olasan")
	}else {
		fmt.Println("do'dir , son kirit ")
	}
}