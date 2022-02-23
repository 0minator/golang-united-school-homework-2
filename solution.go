package square

import "math"

type SidesNum int

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

func CalcSquare(sideLen float64, sidesNum SidesNum) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	case SidesTriangle:
		return math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
	case SidesSquare:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
