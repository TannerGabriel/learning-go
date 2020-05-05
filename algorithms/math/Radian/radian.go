package Radian

import (
	"math"
)

func degreeToRadian(degree float64) float64 {
	return degree * (math.Pi / 180)
}

func radianToDegree(radian float64) float64 {
	return radian * (180 / math.Pi)
}
