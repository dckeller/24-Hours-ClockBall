package main

import (
	"fmt"
	"encoding/json"
	"log"
	"math"
	"reflect"
)

type AllLevels struct {
	Base 		[]int
	Minute 	[]int
	Five		[]int
	Hour		[]int
}

// Create the Base Level with the amount of balls inputed from the user
func (l *AllLevels) MakeBase(quantity int) ([]int, error) {
	l.Base = make([]int, quantity)

	if quantity < 27 || quantity > 127 {
		return nil, fmt.Errorf("%v is not a valid number of balls, must be between 27 - 127.", quantity)
	}
	for i := range l.Base {
		l.Base[i] = i + 1
	}
	return l.Base, nil
}

// Remove the first ball from the Base level
func (l *AllLevels) RemoveBall() {
	l.Base = append(l.Base[:0], l.Base[0+1:]...)
}

// Add removed ball from the Base level to the Minute level
func (l *AllLevels) AddBaseToMin(ball int) {
	l.Minute = append(l.Minute, ball)
}

// Reverse the order of the balls, to place back into the Base level
func ReverseSort(level []int) {

	// i starts at 0, j at the end of slice
	// run until i < j; 
	// i + 1 every loop through, and j - 1 every loop through
	for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {

		// Use a tuple operator to switch places based on index values
		level[i], level[j] == level[j], level[i]
	}
}

// After the balls are reversed, clear the Minute level
func (l *AllLevels) ClearMinBalls() []int {
	l.Minute = append(l.Minute[:0], l.Minute[0+5:]...)
	return l.Minute
}

// Add the fifth ball in the Minute level to the Five level
func (l *AllLevels) AddMinToFive(ball int) {
	l.Five = append(l.Five, ball)
}

// Dump the remaining balls back to the Base level
func (l *AllLevels) DumpBalls(balls []int) []int {
	for _, ball := range balls {
		l.Base = append(l.Base, ball)
	}
	return l.Base
}

// This method adds the Five to Hour level, as well as keeps the 0 ball
// in the Hour level as a placeholder
func (l *AllLevels) AddToHour(ball int) {
	l.Hour = append(l.Hour, ball)
}

// After the balls are reversed, clear the Five level
func (l *AllLevels) ClearFiveBalls() []int {
	l.Five = append(l.Five[:0], l.Five[0+12:]...)
	return l.Five
}

// After the balls are reversed, clear the Hour level
func (l *AllLevels) ClearHourBalls() []int {
	l.Hour = append(l.Hour[:0], l.Hour[0+13:]...)
	return l.Hour
}

// Last ball in the cycle to be added back to the Base level
func (l *AllLevels) AddBackToBase(ball int) {
	l.Base = append(l.Base, ball)
}

// Finishes the program for mode 1
func Finished(a []int, b []int, time float64) bool {
	time = math.Round(time / 60.0 / 24.0)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	fmt.Printf("%v balls finished in %v days \n", len(a), time)
	return true
}

// Finishes the program for mode 2
func (l *AllLevels) ConvertToJson() bool {
	jsonByteData, err := json.Marshal(l)

	if err != nil {
		log.Fatal(err)
		return false
	}
	jsonStringData := string(jsonByteData)
	fmt.Println(jsonStringData)
	return true
}

// The flow of a ball going from the Minute level to the Five level
func (l *AllLevels) LevelMinController() bool {
	var addToFive = l.Minute[4]
	var backToBase = l.Minute[0:4]
	ReverseSort(backToBase)
	l.ClearMinBalls()
	l.AddMinToFive(addToFive)
	l.DumpBalls(backToBase)
	return true
}

// The flow of a ball going from the Five level to the Hour level
func (l *AllLevels) LevelFiveController() bool {
	var addToHour = l.Five[11]
	var backToBase = l.Five[0:11]
	ReverseSort(backToBase)
	l.ClearFiveBalls()
	l.AddToHour(addToHour)
	l.DumpBalls(backToBase)
	return true
}

// The flow of a ball going from the Hour level to the Base level
func (l *AllLevels) LevelHourController() bool {
	var keepInHour = l.Hour[0]
	var lastToBase = l.Hour[12]
	var backToBase = l.Hour[1:12]
	ReverseSort(backToBase)
	l.ClearHourBalls()
	l.AddToHour(keepInHour)
	l.DumpBalls(backToBase)
	l.AddBackToBase(lastToBase)
	return true
}

// Check error function
func CheckError(err error) {
  if err != nil {
    fmt.Println(err)
  }
}

func main() {
	var timer 			float64
	var stop 				bool
	var balls 			int
	var mode				int
	var run 				int
	var floatToInt 	float64

	// Instatitate the All Levels struct, Hour level will always have a 0 placeholder
	all := AllLevels{
		Base: 	make([]int, 0),
		Minute: make([]int, 0),
		Five: 	make([]int, 0),
		Hour: 	[]int{0},
	}

	// Get user input for which mode of the program to run //
	fmt.Println(`Which mode would you like to use?
	1 - Ball Count
	2 - Ball Timer`)
	_, err := fmt.Scanf("%d", &mode)
	CheckError(err)

	if mode == 1 {
		fmt.Println("How many balls would you like to use between 27 - 127?")
		_, err = fmt.Scanf("%d", &balls)
		CheckError(err)

	} else if mode == 2 {
		fmt.Println("How many balls would you like to use between 27 - 127?")
		_, err = fmt.Scanf("%d", &balls)
		CheckError(err)

		fmt.Println("How many minutes would you like to run the program for?")
		_, err = fmt.Scanf("%d", &run)
		CheckError(err)
		floatToInt = float64(run)

	}	else {
		fmt.Println("You must enter a 1 or a 2")
		return
	}

	newBase, err := all.MakeBase(balls)
	if err != nil {
		fmt.Println(err)
		return
	}

	all.Base = make([]int, len(newBase))
	copy(all.Base, newBase)

	for !stop {

		var baseBall = all.Base[0]
		all.RemoveBall()
		all.AddBaseToMin(baseBall)
		timer += 1.0

		if len(all.Minute) == 5 {
			all.LevelMinController()
		}

		if len(all.Five) == 12 {
			all.LevelFiveController()
		}

		if len(all.Hour) == 13 {
			all.LevelHourController()
		}

		if len(newBase) == len(all.Base) && mode == 1 {
			stop = Finished(newBase, all.Base, timer)
		}

		if timer == floatToInt && mode == 2 {
			stop = all.ConvertToJson()
		}
	}
}