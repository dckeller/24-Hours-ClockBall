package main 

import (
	"fmt"
)

// func main() {
// 	num := 210

// 	if num < 200 {
// 		fmt.Println("true")
// 	} else {
// 		fmt.Println("false")
// 	}
// }

// func main() {
// 	num := 11

// 	if num > 5 && num < 10 {
// 		fmt.Println("true")
// 	} else {
// 		fmt.Println("false")
// 	}
// }

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range array {
		fmt.Printf("Let's count nums... %v\n", n)
	}
}