package go_sikuli

import (
	"errors"
	"github.com/go-vgo/robotgo"
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

func performRandomisedClick(double bool) {
	_ = robotgo.Toggle("left", "down")
	sleepRandomly(0.1, 0.2)
	_ = robotgo.Toggle("left", "up")
	if double {
		sleepRandomly(0.2, 0.3)
		_ = robotgo.Toggle("left", "down")
		sleepRandomly(0.1, 0.3)
		_ = robotgo.Toggle("left", "up")
	}
}

// HumanClick simulate human click
func HumanClick(x, y, bW, bH int, double bool) {
	sleepRandomly(0.1, 0.3)
	if bW < 0 {
		bW = 0
	}
	if bH < 0 {
		bH = 0
	}
	moveMouseRandomlyWithinBox(x, y, bW, bH)
	sleepRandomly(0.1, 0.3)
	performRandomisedClick(double)
	_ = robotgo.MouseUp()
}

// ClickImage clicks on the given image within the screen.
// If 'double' is true, it will be a double click. Otherwise, it'll be a single click.
// If an offset is given, it'll click at the x and y offset positions.
// If the image isn't found, an error will be returned.
func ClickImage(imgByte []byte, double bool, humanlike bool, offsets ...int) error {
	img, err := robotgo.ByteToImg(imgByte)
	if err != nil {
		return err
	}
	imgW, imgH := img.Bounds().Dx(), img.Bounds().Dy()

	x, y, err := WaitShow(imgByte, 0.8)
	if err != nil {
		return err
	}
	if x < 0 || y < 0 {
		return errors.New("can't find target")
	}

	if len(offsets) > 0 {
		x = x + offsets[0]
	}
	if len(offsets) > 1 {
		y = y + offsets[1]
	}

	if humanlike {
		HumanClick(x, y, imgW/2-4, imgH/2-4, double)
	} else {
		Click(x, y, double)
	}

	return nil
}

func moveMouseRandomlyWithinBox(x, y, w, h int) {
	randomX := generateRandomNumber(x-w, x+w)
	randomY := generateRandomNumber(y-h, y+h)
	robotgo.MoveSmooth(randomX, randomY, 0.2, 1.1, 1)
}

func sleepRandomly(min, max float64) {
	n := generateRandomNumber(min, max)
	time.Sleep(time.Duration(n) * time.Second)
}
