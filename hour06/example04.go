package main

import "fmt"

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses de Bourgogne"
	fmt.Println(cheeses)
	
	cheeses = append(cheeses, "Camembert", "Reblochon", "Picodon")
	fmt.Println(cheeses)
}
