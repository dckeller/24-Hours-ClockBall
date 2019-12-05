/* recursive function */
package main

import "fmt"

// func feedMe(portion, eaten int) int {
// 	eaten = portion + eaten
// 	if eaten >= 5 {
// 		fmt.Printf("I'm full! I've eaten %d\n", eaten)
// 		return eaten
// 	}
// 	fmt.Printf("I'm still hungry! I've eaten %d\n", eaten)
// 	return feedMe(portion, eaten)
// }

// func main() {
// 	feedMe(1, 0)
// }

func countChocula(count, eaten int) int {
	eaten = count + eaten
	if eaten >= 10 {
		fmt.Printf("Waaha... that's %d chocolates!\n", eaten)
		return eaten
	}
	fmt.Printf("I need more chocolate, I've only had %d today!\n", eaten)
	return countChocula(count, eaten)
}

func main() {
	countChocula(1, 0)
}