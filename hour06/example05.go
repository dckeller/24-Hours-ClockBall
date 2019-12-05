package main

import "fmt"

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses"
	cheeses[2] = "Camembert"
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)

	cheeses = append(cheeses[:1], cheeses[1+2:]...)
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
}
