package main

import (
	"fmt"
	"math"
)

var radius float64 = 4
var centerX float64 = 0
var centerY float64 = 0

var pointX float64 = 1
var pointY float64 = 4

var answer bool = math.Pow(pointX-centerX, 2)+math.Pow(pointY-centerY, 2) < math.Pow(radius, 2)

func main() {
	fmt.Println(answer)
}
