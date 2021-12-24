package termplot

import "math"

func FunctionFromPoints(points [][]float64) func(float64)float64 {
	return func(x float64) float64 {
		for i := 1; i < len(points); i++ {
			if points[i][0] >= x {
				y := (points[i-1][1] * (points[i][0] - x) +
					points[i][1] * (x - points[i-1][0])) / 
						(points[i][0] - points[i-1][0])
				return y
			}
		}
		return math.NaN()
	}
}
