package go_sikuli

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

type IPoint struct {
	X int
	Y int
}

func DePoint(p IPoint) (x, y int) {
	return p.X, p.Y
}

func Find(bigImg []byte, smallImg []byte, threshold float32) IPoint {
	smallMat, _ := gocv.IMDecode(smallImg, gocv.IMReadAnyColor)
	defer smallMat.Close()

	bigMat, _ := gocv.IMDecode(bigImg, gocv.IMReadAnyColor)
	defer bigMat.Close()

	result := gocv.NewMat()
	defer result.Close()

	mask := gocv.NewMat()
	defer mask.Close()
	gocv.MatchTemplate(bigMat, smallMat, &result, gocv.TmCcoeffNormed, mask)

	_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)

	if maxVal < threshold {
		return IPoint{-1, -1}
	}

	w, h, _ := GetImageSize(smallImg)
	p := IPoint{
		X: maxLoc.X + w/2,
		Y: maxLoc.Y + h/2,
	}
	return p
}

func FindAll(bigImg []byte, smallImg []byte, threshold float32) []IPoint {
	smallMat, _ := gocv.IMDecode(smallImg, gocv.IMReadAnyColor)
	defer smallMat.Close()

	bigMat, _ := gocv.IMDecode(bigImg, gocv.IMReadAnyColor)
	defer bigMat.Close()

	result := gocv.NewMat()
	defer result.Close()

	mask := gocv.NewMat()
	defer mask.Close()
	gocv.MatchTemplate(bigMat, smallMat, &result, gocv.TmCcoeffNormed, mask)

	var points []IPoint
	for {
		_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
		if maxVal < threshold {
			break
		}
		w, h, _ := GetImageSize(smallImg)
		points = append(points, IPoint{
			X: maxLoc.X + w/2,
			Y: maxLoc.Y + h/2,
		})
		mask := gocv.NewMatWithSize(smallMat.Rows(), smallMat.Cols(), gocv.MatTypeCV8U)
		mask.SetTo(gocv.NewScalar(0, 0, 0, 0))
		rect := image.Rect(maxLoc.X, maxLoc.Y, maxLoc.X+smallMat.Cols(), maxLoc.Y+smallMat.Rows())
		gocv.Rectangle(&result, rect, color.RGBA{0, 0, 0, 0}, -1)
	}
	return points
}
