package main

import "fmt"

func main(){

	fmt.Println("           Menu              ")
	fmt.Print("1.Go'shtli somsa \n 2.Kartoshkali somsa \n 3.Tovuqli somsa \n 4.Qovoqli somsa \n 5.Tandir Somsa \n 6.Jizzax Somsa \n")

	var menu int
	var bir int
	var ikki int
	var uch int
	var tort int
	var besh int  
	var olti int  

	fmt.Println("        Xoxlagan somsangizni tanlang:       ")
	fmt.Scanf("%d", &menu)
	

	switch menu {
	case 1: 
		fmt.Scan(&bir)
		fmt.Println(bir * 10)
		break
	case 2:
		fmt.Scan( &ikki)
		fmt.Print("Kartoshkali somsa 145 $")
		break
	case 3:
		fmt.Scan(&uch)
		fmt.Print("Tovuqli somsa 160 $")
		break
	case 4:
		fmt.Scan(&tort)
		fmt.Print("Qovoqli somsa 165 $")
		break
    case 5: 
		fmt.Scan(&besh)
		fmt.Print("Tandir Somsa 170 $")
		break
	case 6:
		fmt.Scan(&olti)
		fmt.Print("Jizzax Somsa 175 $")
		break
	default:
		fmt.Println("1 dan 6 gacha son tanlang")
	}
		








}
