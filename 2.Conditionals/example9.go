package main 

import "fmt"

func main(){

number := 0
fmt.Println("Raqam kiriting: ")
fmt.Scanf("%d" , &number)

if number%3 == 0 {
	fmt.Println(number,"Bu son 3 ga bo'linadi")

}else {
	fmt.Println(number,"Bu son 3 ga bo'linmaydi")
}

}