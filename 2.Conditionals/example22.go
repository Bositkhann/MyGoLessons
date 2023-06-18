package main 

import "fmt"

func main(){

	// 			B6
	
	var yil int

	fmt.Println("Ixtiyoriy yilni kiriting :")
	fmt.Scanln(&yil)

	if  yil % 4 == 0 {
		fmt.Println("Bu yil kabisa yil")
	}else if yil % 100 == 0 {
		fmt.Println("Bu yil kabisa yil emas")
	}else {
		fmt.Println("Bu yil kabisa yil emas ")
	}

	




}