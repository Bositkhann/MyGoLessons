package main

import "fmt"

func main(){

	var son int

	fmt.Println("Son kiriting:")
	fmt.Scanln(&son)

	if son == 1 {
		fmt.Println("Bu son Yanvar oyiga tegishli")
	}else if son == 2{
		fmt.Println("Bu son Fevral oyiga tegishli")
	}else if son == 3{
		fmt.Println("Bu son Mart oyiga tegishli")
	}else if son == 4{
		fmt.Println("Bu son Aprel oyiga tegishli")
	}else if son == 5{
		fmt.Println("Bu son May oyiga tegishli")
	}else if son == 6{
		fmt.Println("Bu son Iyun oyiga tegishli")
	}else if son == 7{
		fmt.Println("Bu son Iyul oyiga tegishli")
	}else if son == 8{
		fmt.Println("Bu son Avgust oyiga tegishli")
	}else if son == 9{
		fmt.Println("Bu son Sentyabr oyiga tegishli")
	}else if son == 10{
		fmt.Println("Bu son Oktyabr oyiga tegishli")
	}else if son == 11{
		fmt.Println("Bu son Noyabr oyiga tegishli")
	}else if son == 12{
		fmt.Println("Bu son Dekabr oyiga tegishli")
	}else {
		fmt.Println("Kalla qo'ydin")
	}
}