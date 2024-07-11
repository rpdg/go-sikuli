package go_sikuli

import (
	"math"
	"math/rand"
	"time"
)

type Number interface {
	float64 | float32 | int | int32 | int64
}

func generateRandomNumber[T Number](min T, max T) T {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := (rand.Float64() * (float64(max) - float64(min))) + float64(min)

	// Trims to two decimal places. Doesn't need to be perfect.
	return T(math.Floor(float64(randNum)*100) / 100)
}

// abs is a helper function that returns the absolute value of an integer
func abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
