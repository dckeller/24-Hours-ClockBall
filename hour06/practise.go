package main 

import (
	"fmt"
)

// func main() {
// 	var sports = make([]string, 4)

// 	sports[0] = "baseball"
// 	sports[1] = "basketball"
// 	sports[2] = "golf"
// 	sports[3] = "swimming"

// 	fmt.Println(sports)

// 	var balls = make([]string, 2)
// 	copy(balls, sports[0:2])

// 	fmt.Println(balls)
// }

func main() {
	var elements = make(map[string]string)

	elements["p"] = "Paragraph"
	elements["img"] = "Image"
	elements["h1"] = "Heading One"
	elements["h2"] = "Heading Two"

	fmt.Println(elements)
}
