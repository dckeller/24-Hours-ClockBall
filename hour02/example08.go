package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var b bool = true
	fmt.Println(reflect.TypeOf(b))
	
	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
	
	var t string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(t))

	n := strconv.Itoa(123)
	fmt.Println(n)
	fmt.Println(reflect.TypeOf(n))

	m, _ := strconv.Atoi("42")
	fmt.Println(reflect.TypeOf(m))
}
