package go_sikuli

import (
	"github.com/go-vgo/robotgo"
	"testing"
)

func Test_FindImg(t *testing.T) {
	t.Run("test find one", func(t *testing.T) {
		big, _ := robotgo.OpenImg("C:\\tests\\big.png")
		small, _ := robotgo.OpenImg("C:\\tests\\small.png")
		p := Find(big, small, 0.9)
		println(p.X, p.Y)
	})

	t.Run("test find all", func(t *testing.T) {
		big, _ := robotgo.OpenImg("C:\\tests\\big.png")
		small, _ := robotgo.OpenImg("C:\\tests\\small.png")
		ps := FindAll(big, small, 0.9)
		for i, p := range ps {
			println("point-", i, p.X, p.Y)
		}
	})
}
