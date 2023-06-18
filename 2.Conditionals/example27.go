package main
	
import "fmt"
	
func main() {

	fmt.Println("     Restoran Menyusi")
	fmt.Println("1. Shashlik - 15000 so'm")
	fmt.Println("2. Osh 2 kg - 100000 so'm")
	fmt.Println("3. Lag'mon - 22000 so'm")
	fmt.Println("4. Chuchvara - 28000 so'm")
	
	
	var jami int
	var menu int


	fmt.Scan(&menu)
	
	switch menu {
	case 1:
		fmt.Println("Shashlik  15000 som")
		jami += 15000
		
	case 2:
		fmt.Println("Osh 2kg 100000 som")
		jami += 100000
		
	case 3:
		fmt.Println("Lag'mon 22000 som")
		jami += 22000
		
	case 4:
		fmt.Println("Chuchvara  28000 som")
		jami += 28000 
	
	break
	

	default:
		fmt.Println("Xato ")
	}
	 
	
	fmt.Printf("Jami: %d som", jami)
}