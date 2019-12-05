package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}

// func main() {
// 	m := Movie{
// 		Name:   "Citizen Kane",
// 		Rating: 10,
// 	}
// 	fmt.Printf("%+v\n", m)
// }

// func main() {
// 	var m Movie

// 	m.Name = "The Sandlot"
// 	m.Rating = 7.6

// 	fmt.Printf("%+v\n", m)
// }

// func main() {
// 	m := new(Movie)

// 	m.Name = "The Moulin Rouge"
// 	m.Rating = 9.9

// 	fmt.Printf("%+v\n", m)
// }