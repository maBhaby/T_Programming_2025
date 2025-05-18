package main

import (
	"fmt"
	"math"
)

func calculateY(x float64) (float64, error) {
	if x < -1 || x > 1 {
		return 0, fmt.Errorf("x must be in the range [-1, 1]")
	}

	arcsinX := math.Asin(x)
	arccosX := math.Acos(x)

	arcsin4 := math.Pow(arcsinX, 4)
	arccos6 := math.Pow(arccosX, 6)

	sum := arcsin4 + arccos6
	y := math.Cbrt(sum)

	return y, nil
}

func main() {
	x := 0.5

	y, err := calculateY(x)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("При x = %.2f, y = %.6f\n", x, y)
}
