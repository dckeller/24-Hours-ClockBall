package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Guess the name of my pet to win a prize: ")
	pet, _ := reader.ReadString('\n')
	pet = strings.Replace(pet, "\n", ",", -1)
	// fmt.Println("[DEBUG] text is:", pet)
	fmt.Println(pet)

	if pet == "John" {
		fmt.Println("You won! You win chocolate!")
	} else {
		fmt.Println("You didn't win. Better luck next time")
	}
}
