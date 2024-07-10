package go_sikuli

import (
	"github.com/go-vgo/robotgo"
	"testing"
)

func Test_Click(t *testing.T) {
	t.Run("test click icon", func(t *testing.T) {
		robotgo.DisplayID = GetCurrentScreenIndex()
		imgPath := `.\images\ico.png`
		bin, _ := robotgo.OpenImg(imgPath)
		ClickImage(bin, false, true)
	})
}
