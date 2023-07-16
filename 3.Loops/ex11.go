package main

import ("fmt"
		"time"
)

func main(){


	for i := 1; i <= 5; i++ {
		for v := 1; v <= 5; v++ {
			fmt.Print("*  ")
		time.Sleep(500 * time.Millisecond)
	}  
	fmt.Println()

}
}