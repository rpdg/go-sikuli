package go_sikuli

import (
	"errors"
	"github.com/go-vgo/robotgo"
	"math"
	"math/rand"
	"time"
)

// Click performs a left mouse click at the given x and y coordinates.
// If double is set to true, it will perform a double-click.
// A random sleep time is added in between mouse down and mouse up actions
// to ensure the actions are performed correctly.
func Click(x, y int, double bool) {
	robotgo.Move(x, y)
	sleepRandomly(0.2, 0.5)
	robotgo.Click("left", double)
	_ = robotgo.MouseUp()
	sleepRandomly(0.2, 0.5)
}

// ClickImage clicks on the given image within the screen.
// If 'double' is true, it will be a double click. Otherwise it'll be a single click.
// If an offset is given, it'll click at the x and y offset positions.
// If the image isn't found, an error will be returned.
func ClickImage(imgByte []byte, double bool, offset ...int) error {
	x, y, err := WaitShow(imgByte, 0.88)
	if err != nil {
		return err
	}

	if x < 0 || y < 0 {
		return errors.New("cant find target")
	}

	if len(offset) > 0 {
		x = x + offset[0]
		if len(offset) > 1 {
			y = y + offset[1]
		}
	}
	Click(x, y, double)
	return nil
}

// HumanClick simulate human click
func HumanClick(x, y int, double bool) {
	sleepRandomly(0.2, 0.5)
	moveMouseRandomlyWithinBox(float64(x), float64(y), 3, 3)
	sleepRandomly(0.2, 0.5)
	performRandomisedClick(double)
	_ = robotgo.MouseUp()
}

func performRandomisedClick(double bool) {
	_ = robotgo.Toggle("left", "down")
	sleepRandomly(0.1, 0.3)
	_ = robotgo.Toggle("left", "up")
	if double {
		sleepRandomly(0.2, 0.4)
		_ = robotgo.Toggle("left", "down")
		sleepRandomly(0.1, 0.3)
		_ = robotgo.Toggle("left", "up")
	}
}
func moveMouseRandomlyWithinBox(x, y, w, h float64) {
	randomX := generateRandomNumber(x, x+w)
	randomY := generateRandomNumber(y, y+h)

	robotgo.Move(int(math.Round(randomX)), int(math.Round(randomY)))
}
func sleepRandomly(min, max float64) {
	n := generateRandomNumber(min, max)
	time.Sleep(time.Duration(n) * time.Second)
}
func generateRandomNumber(min float64, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	randNum := (rand.Float64() * (max - min)) + min

	// Trims to two decimal places. Doesn't need to be perfect.
	return math.Floor(randNum*100) / 100
}
