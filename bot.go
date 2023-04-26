package go_sikuli

import (
	"errors"
	"github.com/go-vgo/robotgo"
)

func ClickPoint(p IPoint) {
	ClickXY(p.X, p.Y)
}

func ClickXY(x, y int) {
	robotgo.Move(x, y)
	robotgo.MilliSleep(200)
	robotgo.Click("left", false)
	_ = robotgo.MouseUp()
}

// ClickImage clicks on screen according inputted image byte
func ClickImage(imgByte []byte, offset ...int) error {
	x, y, err := WaitShow(imgByte)
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
	ClickXY(x, y)
	return nil
}

func KeyTap(key string, args ...interface{}) error {
	return robotgo.KeyTap(key, args...)
}

func TypeStr(str string, args ...int) {
	robotgo.TypeStr(str, args...)
}

func KeyPress(str string, args ...interface{}) error {
	return robotgo.KeyPress(str, args...)
}
