package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"reflect"
	"github.com/pkg/profile"
)

type AllLevels struct {
	Base   []int
	Minute []int
	Five   []int
	Hour   []int
}

func (l *AllLevels) AddBaseToMin(ball int) {
	l.Minute = append(l.Minute, ball)
}

func (l *AllLevels) AddMinToFive(ball int) {
	l.Five = append(l.Five, ball)
}

func (l *AllLevels) AddToHour(ball int) {
	l.Hour = append(l.Hour, ball)
}

func (l *AllLevels) AddBackToBase(ball int) {
	l.Base = append(l.Base, ball)
}

func (l *AllLevels) ClearMinBalls() []int {
	l.Minute = append(l.Minute[:0], l.Minute[0+5:]...)
	return l.Minute
}

func (l *AllLevels) ClearFiveBalls() []int {
	l.Five = append(l.Five[:0], l.Five[0+12:]...)
	return l.Five
}

func (l *AllLevels) ClearHourBalls() []int {
	l.Hour = append(l.Hour[:0], l.Hour[0+13:]...)
	return l.Hour
}

func ReverseSort(level []int) {
	s := reflect.ValueOf(level)
	swp := reflect.Swapper(s.Interface())
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		swp(i, j)
	}
}

func (l *AllLevels) DumpBalls(balls []int) []int {
	for _, ball := range balls {
		l.Base = append(l.Base, ball)
	}
	return l.Base
}

func (l *AllLevels) RemoveBall() {
	l.Base = append(l.Base[:0], l.Base[0+1:]...)
}

func (l *AllLevels) MakeBase(quantity int) ([]int, error) {
	l.Base = make([]int, quantity)

	if quantity < 27 || quantity > 127 {
		return nil, fmt.Errorf("%v is not a valid number, must be 27 - 127", quantity)
	}
	for i := range l.Base {
		l.Base[i] = 1 + i
	}
	return l.Base, nil
}

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

func (l *AllLevels) LevelMinController() bool {
	var addToFive = l.Minute[4]
	var backToBase = l.Minute[0:4]
	ReverseSort(backToBase)
	l.ClearMinBalls()
	l.AddMinToFive(addToFive)
	l.DumpBalls(backToBase)
	return true
}

func (l *AllLevels) LevelFiveController() bool {
	var addToHour = l.Five[11]
	var backToBase = l.Five[0:11]
	ReverseSort(backToBase)
	l.ClearFiveBalls()
	l.AddToHour(addToHour)
	l.DumpBalls(backToBase)
	return true
}

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

func main() {
	// CPU profiling by default
	defer profile.Start().Stop()
	
	timer := 0.0

	// instantiate levels struct
	all := AllLevels{
		Base:   make([]int, 0),
		Minute: make([]int, 0),
		Five:   make([]int, 0),
		Hour:   []int{0},
	}

	// get user input for ball count
	fmt.Println("How many balls would you like to use?")
	var Balls int
	_, err := fmt.Scanf("%d", &Balls)

	// get user input for minutes to run
	// fmt.Println("How many minutes would you like to go for?")
	// var FloatToInt float64
	// var mins int
	// _, err = fmt.Scanf("%d", &mins)
	// FloatToInt = float64(mins)

	// create base level
	newBase, err := all.MakeBase(Balls)
	if err != nil {
		fmt.Println(err)
		return
	}
	all.Base = make([]int, len(newBase))

	// create a copy of the base to stop
	copy(all.Base, newBase)

	stop := false

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

		if len(newBase) == len(all.Base) {
			stop = Finished(newBase, all.Base, timer)
		}

		// if timer == FloatToInt {
		// 	stop = all.ConvertToJson()
		// }
	}
}
