package go_sikuli

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"testing"
)

func Test_Click(t *testing.T) {
	t.Run("test click icon", func(t *testing.T) {
		bin, _ := robotgo.OpenImg(`.\images\ico.png`)
		n := robotgo.DisplaysNum()
		fmt.Println(n)
		robotgo.DisplayID = GetCurrentScreenIndex()
		ClickImage(bin, true, true)
	})
}
