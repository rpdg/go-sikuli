package go_sikuli

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"log"
	"os"
	"testing"
)

func Test_Click(t *testing.T) {
	//robotgo.DisplayID = GetCurrentScreenIndex()
	t.Run("test click icon", func(t *testing.T) {
		imgPath := `.\images\ico.png`
		bin, _ := robotgo.OpenImg(imgPath)
		ClickImage(bin, false, true)
	})
}

func Test_Find(t *testing.T) {
	robotgo.DisplayID = GetCurrentScreenIndex()
	s := ScreenShotByte()
	err := os.WriteFile(`.\images\snap.png`, s, 0644)
	if err != nil {
		log.Fatal(err)
	}
	t.Run("test find all", func(t *testing.T) {
		imgPath := `.\images\ico.png`
		bin, _ := robotgo.OpenImg(imgPath)
		p := FindAll(s, bin, 0.8)
		fmt.Println(p)
	})
}
