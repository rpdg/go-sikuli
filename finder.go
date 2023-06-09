package go_sikuli

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

// Find takes two images and a threshold as parameters and
// finds the location of the small image within the larger image.
// It returns a point which is the center of the small image in the larger image
// if it can be found, or {-1,-1} if it cannot.
func Find(bigImg []byte, smallImg []byte, threshold float32) *image.Point {
	w, h, err := GetImageSize(smallImg)
	if err != nil {
		return nil
	}

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
		return nil
	}

	p := &image.Point{
		X: maxLoc.X + w/2,
		Y: maxLoc.Y + h/2,
	}
	return p
}

// FindAll is a function used to find all the instances of a small image
// within a big image, up to a given threshold. It takes two byte slices for
// the big and small images, and a float value for the threshold, and returns
// a slice of image.Point instances for each instance found.
func FindAll(bigImg []byte, smallImg []byte, threshold float32) []image.Point {
	w, h, err := GetImageSize(smallImg)
	if err != nil {
		return nil
	}

	smallMat, _ := gocv.IMDecode(smallImg, gocv.IMReadAnyColor)
	defer smallMat.Close()

	bigMat, _ := gocv.IMDecode(bigImg, gocv.IMReadAnyColor)
	defer bigMat.Close()

	result := gocv.NewMat()
	defer result.Close()

	mask := gocv.NewMat()
	defer mask.Close()
	gocv.MatchTemplate(bigMat, smallMat, &result, gocv.TmCcoeffNormed, mask)

	var points []image.Point
	for {
		_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
		if maxVal < threshold {
			break
		}

		points = append(points, image.Point{
			X: maxLoc.X + w/2,
			Y: maxLoc.Y + h/2,
		})

		// 在结果矩阵上绘制一个填充矩形，以便稍后不会再次检测到该匹配项
		mask.SetTo(gocv.NewScalar(0, 0, 0, 0))
		rect := image.Rect(maxLoc.X, maxLoc.Y, maxLoc.X+smallMat.Cols(), maxLoc.Y+smallMat.Rows())
		gocv.Rectangle(&result, rect, color.RGBA{}, -1)
	}

	return filterPoints(points)
}

// filterPoints is a helper function that filters an array of image.Point by
// removing any points which are within a 5 pixel distance from any other points.
// It returns an array of distinct, filtered points.
func filterPoints(points []image.Point) []image.Point {
	filtered := make([]image.Point, 0, len(points))
	for _, p1 := range points {
		include := false
		for j := 0; j < len(filtered); j++ {
			p2 := filtered[j]
			if abs(p1.X-p2.X) <= 5 && abs(p1.Y-p2.Y) <= 5 {
				include = true
				break
			}
		}
		if !include {
			filtered = append(filtered, p1)
		}
	}
	return filtered
}

// abs is a helper function that returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
