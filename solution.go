package square

import (
	"math"
)

type intCustomType int

var SidesTriangle intCustomType = 3
var SidesSquare intCustomType = 4
var SidesCircle intCustomType = 0

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	var square float64
	if sidesNum == 3 {
		square = (math.Sqrt(3) / 4) * math.Pow(sideLen, 2)
	} else if sidesNum == 4 {
		square = math.Pow(sideLen, 2)
	} else if sidesNum == 0 {
		square = math.Pi * math.Pow(sideLen, 2)
	} else {
		square = 0
	}
	return square
}
