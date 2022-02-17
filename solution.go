package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sidesNumber int

const SidesCircle sidesNumber = 0
const SidesTriangle sidesNumber = 3
const SidesSquare sidesNumber = 4

func CalcSquare(sideLen float64, sidesNum sidesNumber) float64 {
	return sidesNum.CalcSquare(sideLen)
}

func (s sidesNumber) CalcSquare(sideLen float64) float64 {
	switch s {
	case SidesCircle:
		return circleArea(sideLen)
	case SidesTriangle:
		return triangleArea(sideLen)
	case SidesSquare:
		return squareArea(sideLen)
	default:
		return 0
	}

}

func triangleArea(sideLen float64) float64 {
	return math.Pow(2, sideLen) * math.Sqrt(3) / 4
}

func squareArea(sideLen float64) float64 {
	return math.Pow(2, sideLen)

}

func circleArea(sideLen float64) float64 {
	return math.Pi * math.Pow(2, sideLen)
}
