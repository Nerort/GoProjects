package main

import (
	"fmt"
	"math"
)

var aSide float64 = 5
var bSide float64 = 4
var hypotenuse float64 = math.Sqrt(math.Pow(aSide, 2) + math.Pow(bSide, 2))

func main() {
	fmt.Println(hypotenuse)
}
