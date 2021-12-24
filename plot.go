package termplot

import "math"

func FunctionsMinAndMax(xMin, xMax float64, functions ...func(float64)float64) (float64, float64) {
	step := (xMax - xMin) / 100
	yMin := math.Inf(1)
	yMax := math.Inf(-1)
	for x := xMin; x <= xMax; x+=step {
		for _, function := range functions {
			value := function(x)
			if value < yMin {
				yMin = value
			}
			if value > yMax {
				yMax = value
			}
		}
	}
	return yMin, yMax
}

func PointsFromFunction(function func(float64) float64, width, height int, xMin, xMax, yMin, yMax float64) (points [][]int) {
	values := calculateValues(function, xMin, xMax, width)
	values = normalize(values, height, yMin, yMax)

	lastY := -1
	for x, value := range values {
		if math.IsNaN(value) {
			lastY = -1
			continue
		}
		y := int(value)
		points = append(points, []int{x, y})

		if lastY != -1 && y != lastY {
			step := (lastY - y) / abs(lastY - y)
			for traceY := y; traceY != lastY; traceY += step {
				points = append(points, []int{x, traceY})
			}
		}
		lastY = y
	}

	return points
}

func calculateValues(function func(float64) float64, xMin, xMax float64, width int) (values []float64) {
	var xStep = (xMax - xMin) / float64(width)
	for i := 0; i < width; i++ {
		values = append(values, function(float64(i) * xStep + xMin))
	}
	return values
}

func normalize(values []float64, height int, yMin, yMax float64) (normalized []float64) {
	yScale := float64(height) / (yMax - yMin)

	for _, value := range values {
		normalized = append(
			normalized,
			bound((value - yMin) * yScale, 0, float64(height-1)),
		)
	}
	
	return normalized
}

type Color string
var Reset Color = "\033[0m"
var White Color = "\033[47m"
var Black Color = "\033[46m"
var Red Color = "\033[41m"
var Green Color = "\033[42m"
var Yellow Color = "\033[43m"
var Blue Color = "\033[44m"
var Magenta Color = "\033[45m"
var Cyan Color = "\033[46m"

func DrawPoints(plots [][][]int, colors []Color, width, height int) string {
	graph := generateCanvas(width, height)

	for i := range plots {
		for _, point := range plots[i] {
			graph[point[1]][point[0]] = string(colors[i]) + " " + string(Reset)
		}

	}

	printString := ""
	for i:=len(graph)-1; i>=0; i-- {
		for _, str := range graph[i] {
			printString += str
		}
		printString += "\n"
	}

	return printString
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func bound(value, min, max float64) float64 {
	if value <= min {
		return math.NaN()
	} else if value >= max {
		return math.NaN()
	} else {
		return value
	}
}

func min(slice []float64) (result float64) {
	for _, value := range slice {
		if value < result {
			result = value
		}
	}
	return result
}

func max(slice []float64) (result float64) {
	for _, value := range slice {
		if value > result {
			result = value
		}
	}
	return result
}

func generateCanvas(width, height int) (canvas [][]string) {
	for i := 0; i < height; i++ {
		var line []string
		for j := 0; j < width; j++ {
			line = append(line, " ")
		}
		canvas = append(canvas, line)
	}
	return canvas
}
