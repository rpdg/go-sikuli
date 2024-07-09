package go_sikuli

import (
	"encoding/json"
	"fmt"
	"github.com/go-vgo/robotgo"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"testing"
)

func printStruct(i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func TestHumanDetect(t *testing.T) {
	// define default hog descriptor
	hog := gocv.NewHOGDescriptor()
	defer hog.Close()
	hog.SetSVMDetector(gocv.HOGDefaultPeopleDetector())

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// read image
	img := gocv.IMRead("images/2boFZ.png", gocv.IMReadGrayScale)

	//resize image
	fact := float64(400) / float64(img.Cols())
	newY := float64(img.Rows()) * fact
	gocv.Resize(img, &img, image.Point{X: 400, Y: int(newY)}, 0, 0, 1)

	// detect people in image
	rects := hog.DetectMultiScaleWithParams(img, 0, image.Point{X: 8, Y: 8}, image.Point{X: 16, Y: 16}, 1.05, 2, false)

	// print found points
	printStruct(rects)

	// draw a rectangle around each face on the original image,
	// along with text identifing as "Human"
	for _, r := range rects {
		gocv.Rectangle(&img, r, blue, 3)

		size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
		pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
		gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
	}

	if ok := gocv.IMWrite("loool.jpg", img); !ok {
		fmt.Println("Error")
	}
}

func Test_FindImg(t *testing.T) {
	t.Run("test find FindSIFT", func(t *testing.T) {
		big, _ := robotgo.OpenImg("C:\\tests\\big.png")
		small, _ := robotgo.OpenImg("C:\\tests\\small.png")
		p, err := FindSIFT(small, big)
		if err != nil {
			t.Error(err)
		} else {
			printStruct(p)
		}
	})

	t.Run("test find one", func(t *testing.T) {
		big, _ := robotgo.OpenImg("C:\\tests\\big.png")
		small, _ := robotgo.OpenImg("C:\\tests\\small.png")
		p := Find(big, small, 0.9)
		if p != nil {
			printStruct(p)
		}
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
