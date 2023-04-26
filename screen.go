package go_sikuli

import (
	"errors"
	"github.com/go-vgo/robotgo"
	"gocv.io/x/gocv"
	"image"
	"time"
)

// ShotByte use gocv to capture the screen and return it as byte array
func ShotByte() []byte {
	window := gocv.NewWindow("Capture Screen")
	defer window.Close()
	img := gocv.IMRead("0", gocv.IMReadColor)
	defer img.Close()

	if img.Empty() {
		return nil
	}

	buf, _ := gocv.IMEncode(".png", img)
	return buf.GetBytes()
}

func Snap() image.Image {
	screenBmp := robotgo.CaptureScreen()
	defer robotgo.FreeBitmap(screenBmp)
	time.Sleep(time.Millisecond * 200)
	return robotgo.ToImage(screenBmp)
}

func SnapByte() []byte {
	scrImg := Snap()
	return ImageToByte(scrImg)
}

func WaitShow(target []byte, timeouts ...int) (int, int, error) {
	maxSecond := 30
	if len(timeouts) > 0 {
		maxSecond = timeouts[0]
	}
	i := 0
	if len(timeouts) > 1 {
		i = timeouts[1]
	}
	scrByt := SnapByte()
	p := Find(scrByt, target, 0.7)
	if p.X < 0 || p.Y < 0 {
		i++
		if i < maxSecond*4 {
			time.Sleep(time.Millisecond * 250)
			return WaitShow(target, maxSecond, i)
		} else {
			return -1, -1, errors.New("timeout")
		}
	}
	x, y := DePoint(p)
	return x, y, nil
}

func WaitHide(target []byte, timeouts ...int) error {
	maxSecond := 30
	if len(timeouts) > 0 {
		maxSecond = timeouts[0]
	}
	i := 0
	if len(timeouts) > 1 {
		i = timeouts[1]
	}
	scrByt := SnapByte()
	x, y := DePoint(Find(scrByt, target, 0.7))
	if x >= 0 && y >= 0 {
		i++
		if i < maxSecond*4 {
			time.Sleep(time.Millisecond * 250)
			return WaitHide(target, maxSecond, i)
		}
		return errors.New("timeout")
	}
	return nil
}
