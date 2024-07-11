package go_sikuli

import (
	"errors"
	"github.com/go-vgo/robotgo"
	"time"
)

func ScreenShotByte() []byte {
	scrImg := robotgo.CaptureImg()
	return ImageToByte(scrImg)
}

func WaitShow(target []byte, threshold float32, timeouts ...int) (x, y int, err error) {
	maxSecond := 30
	if len(timeouts) > 0 {
		maxSecond = timeouts[0]
	}
	i := 0
	if len(timeouts) > 1 {
		i = timeouts[1]
	}
	scrByt := ScreenShotByte()
	p := Find(scrByt, target, threshold)
	if p == nil {
		i++
		if i < maxSecond*4 {
			time.Sleep(time.Millisecond * 250)
			return WaitShow(target, threshold, maxSecond, i)
		} else {
			return -1, -1, errors.New("timeout")
		}
	}
	return p.X, p.Y, nil
}

func WaitHide(target []byte, threshold float32, timeouts ...int) error {
	maxSecond := 30
	if len(timeouts) > 0 {
		maxSecond = timeouts[0]
	}
	i := 0
	if len(timeouts) > 1 {
		i = timeouts[1]
	}
	scrByt := ScreenShotByte()
	p := Find(scrByt, target, threshold)
	if p != nil {
		i++
		if i < maxSecond*4 {
			time.Sleep(time.Millisecond * 250)
			return WaitHide(target, threshold, maxSecond, i)
		}
		return errors.New("timeout")
	}
	return nil
}
