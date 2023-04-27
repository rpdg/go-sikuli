package go_sikuli

import (
	"github.com/go-vgo/robotgo"
	"testing"
)

func Test_Click(t *testing.T) {
	t.Run("test click icon", func(t *testing.T) {
		bin, _ := robotgo.OpenImg("C:\\tests\\bin.png")
		ClickImage(bin, true)
	})
}
