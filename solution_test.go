package square

import "testing"

func TestCalcSquare(t *testing.T) {
	var circle = 7.0685834705770345
	if got := CalcSquare(1.5, SidesCircle); got != circle {
		t.Errorf("CalcSquare() = %v, want %v", got, circle)
	}

	var triangle = 0.9742785792574934
	if got := CalcSquare(1.5, SidesTriangle); got != triangle {
		t.Errorf("CalcSquare() = %v, want %v", got, triangle)
	}

	var square = 2.25
	if got := CalcSquare(1.5, SidesSquare); got != square {
		t.Errorf("CalcSquare() = %v, want %v", got, square)
	}
}
