package go_sikuli

import (
	"fmt"
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

// FindSIFT takes two images as parameters and finds the location of the small image
// within the larger image using SIFT feature matching.
// It returns a point which is the center of the small image in the larger image
// if it can be found, or {-1,-1} if it cannot.
func FindSIFT(bigImg []byte, smallImg []byte) (*image.Point, error) {
	w, h, err := GetImageSize(smallImg)
	if err != nil {
		return nil, fmt.Errorf("failed to get small image size: %w", err)
	}

	smallMat, _ := gocv.IMDecode(smallImg, gocv.IMReadAnyColor)
	defer smallMat.Close()

	bigMat, _ := gocv.IMDecode(bigImg, gocv.IMReadAnyColor)
	defer bigMat.Close()

	// 初始化gocv的SIFT检测器，并用它来找到图像中的关键点和描述符
	sift := gocv.NewSIFT()
	defer sift.Close()
	_, descriptors1 := sift.DetectAndCompute(smallMat, gocv.NewMat())
	keypoints2, descriptors2 := sift.DetectAndCompute(bigMat, gocv.NewMat())

	// 创建一个BFMatcher对象，使用gocv.NormL2作为距离度量
	matcher := gocv.NewBFMatcherWithParams(gocv.NormL2, false)
	defer matcher.Close()

	// 使用BFMatcher.KnnMatch()方法来匹配两个图像中的描述符，并根据距离排序
	matches := matcher.KnnMatch(descriptors1, descriptors2, 2)

	// 使用阈值来过滤掉不好的匹配，或者使用Lowe的比率测试来过滤掉不好的匹配
	var goodMatches []gocv.DMatch
	for _, m := range matches {
		if len(m) >= 2 {
			if m[0].Distance < 0.75*m[1].Distance {
				goodMatches = append(goodMatches, m[0])
			}
		}

	}

	// 返回匹配点对中的第一个点作为image.Point
	if len(goodMatches) > 0 {
		p := keypoints2[goodMatches[0].QueryIdx]
		return &image.Point{X: int(p.X) + w/2, Y: int(p.Y) + h/2}, nil
	} else {
		return nil, fmt.Errorf("not enough good matches")
	}
}
