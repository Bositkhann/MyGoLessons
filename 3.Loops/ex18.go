package main

import (
	"fmt"
	"time"
)

func main(){

var a int = 0  // SEKUND
var b int = 0  // MINUT
var c int = 0  // Soat
var d int = 0  // Kun
var e int = 0  // Oy
var f int = 0  // Yil

for true {
a++



time.Sleep(1 * time.Second)
fmt.Printf("Sekund: %d  Minut: %d  Soat: %d  Kun: %d  Oy: %d  Yil: %d \n",a,b,c,d,e,f)
if a >= 60 {
	a = 0
	b += 1
}else if b >= 60 {
	b = 0
	c += 1
}else if c >= 24 {
	c = 0 
	d += 1
}else if d >= 30 {
	d = 0 
	e += 1
}else if e >= 12 {
	e = 0 
	f += 1
}
}
}