package main

import (
	"fmt"
	"reflect"
	"math"
	// "encoding/json"
	// "log"
)

type baseLevel struct {
	base   	[]int 	`json:"Main"`
}

type minLevel struct {
	minute 	[]int 	`json:"Min"`
}

type fiveLevel struct {
	five 		[]int 	`json:"FiveMin"`
}

type hourLevel struct {
	hour 		[]int 	`json:"Hour"`
}

func (b *baseLevel) addBall(ball int) []int {
	b.base = append(b.base, ball)
	return b.base
}

func (m *minLevel) addBall(ball int) []int {
	m.minute = append(m.minute, ball)
	return m.minute
}

func (f *fiveLevel) addBall(ball int) []int {
	f.five = append(f.five, ball)
	return f.five
}

func (h *hourLevel) addBall(ball int) []int {
	h.hour = append(h.hour, ball)
	return h.hour
}


func (m *minLevel) clearBalls() []int {
	m.minute = append(m.minute[:0], m.minute[0+5:]...)
	return m.minute
}

func (f *fiveLevel) clearBalls() []int {
	f.five = append(f.five[:0], f.five[0+12:]...)
	return f.five
}

func (h *hourLevel) clearBalls() []int {
	h.hour = append(h.hour[:0], h.hour[0+13:]...)
	return h.hour
}

func reverseSort(level []int) {
	s := reflect.ValueOf(level)
	swp := reflect.Swapper(s.Interface())
	for i, j := 0, s.Len() - 1; i<j; i,j = i+1, j-1 {
		swp(i,j)
	}
}

func (b *baseLevel) dumpBalls(balls []int) []int {
	for _, ball := range balls {
		b.base = append(b.base, ball)
	}
	return b.base
}

func (b *baseLevel) removeBall() []int {
	b.base = append(b.base[:0], b.base[0+1:]...)
	return b.base
}

func (b *baseLevel) makeBase(quantity int) ([]int, error) {
	base := make([]int, quantity)

	if quantity < 27 || quantity > 127 {
		return nil, fmt.Errorf("%v is not a valid number, must be 27 - 127", quantity)
	}
	for i := range base {
		base[i] = 1 + i
	}
	return base, nil
}

func finished(a []int, b []int, time float64) bool {
	time = math.Round(time/60.0/24.0)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	fmt.Printf("%v balls finished in %v days \n", len(a), time)
	return true 
}

// func (b *baseLevel) convertToJson() {
// 	jsonByteData, err := json.Marshal(b)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	jsonStringData := string(jsonByteData)
// }

// func (m *minLevel) convertToJson() {
// 	jsonByteData, err := json.Marshal(m)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	jsonStringData := string(jsonByteData)
// }

// func (f *fiveLevel) convertToJson() {
// 	jsonByteData, err := json.Marshal(f)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	jsonStringData := string(jsonByteData)
// }

// func (h *hourLevel) convertToJson() {
// 	jsonByteData, err := json.Marshal(h)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	jsonStringData := string(jsonByteData)
// }


func main() {
	timer := 0.0

	var b baseLevel

	var m minLevel
	m.minute = make([]int, 0)

	var f fiveLevel
	f.five = make([]int, 0)

	var h hourLevel
	h.hour = []int{0}

	fmt.Println("How many balls would you like to use?")
	var balls int
	_, err := fmt.Scanf("%d", &balls)

	// fmt.Println("How many minutes would you like to run?")
	// var mins float64
	// _, err := fmt.Scanf("%d", &mins)


	//start with inputed base//
	Base, err := b.makeBase(balls)
	if err != nil {
		fmt.Println(err)
		return
	}

	b.base = make([]int, len(Base))

	copy(b.base, Base)
	
	stop := false

	for !stop {	

		var addToMin = b.base[0]
		b.removeBall()
		m.addBall(addToMin)
		timer += 1.0 

		if len(m.minute) == 5 {
			var addToFive = m.minute[4]
			var backToBase = m.minute[0:4]
			reverseSort(backToBase)
			m.clearBalls()
			f.addBall(addToFive)
			b.dumpBalls(backToBase)
		}

		if len(f.five) == 12 {
			var addToHour = f.five[11]
			var backToBase = f.five[0:11]
			reverseSort(backToBase)
			f.clearBalls()
			h.addBall(addToHour)
			b.dumpBalls(backToBase)
		}

		if len(h.hour) == 13 {
			var keepInHour = h.hour[0]
			var lastToBase = h.hour[12]
			var backToBase = h.hour[1:12]
			reverseSort(backToBase)
			h.clearBalls()
			h.addBall(keepInHour)
			b.dumpBalls(backToBase)
			b.addBall(lastToBase)
		}

		if len(Base) == len(b.base) {
			stop = finished(Base, b.base, timer)
		}

		// if timer == mins {
		// 	stop = 
		// }

		fmt.Printf("Hour: %v\n", h.hour)
		fmt.Printf("Five: %v\n", f.five)
		fmt.Printf("Min: %v\n", m.minute)
		fmt.Printf("Base: %v\n", b.base)
		fmt.Println("\n")

	}	
}	