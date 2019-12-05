package main 

import (
	"fmt"
)

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

// Moving the c to either side will result in using the vals() or a or b
	c, _ := vals() 
	fmt.Println(c)	
}
