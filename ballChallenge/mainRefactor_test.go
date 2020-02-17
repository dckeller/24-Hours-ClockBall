package main

import "testing"


func BenchmarkAddBaseToMin(b *testing.B) {
	AddBaseToMin(1)
}