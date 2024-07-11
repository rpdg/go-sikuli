package go_sikuli

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/png"
	"regexp"
)

func GetImageSize(imgBytes []byte) (int, int, error) {
	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return 0, 0, err
	}
	bounds := img.Bounds()
	return bounds.Dx(), bounds.Dy(), nil
}

func ImageToByte(img image.Image) []byte {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func DecodeBase64Image(base64Str string) (imgBytes []byte, err error) {
	re := regexp.MustCompile(`^data:image/(png|jpg);base64,`)
	trimmedBase64Str := re.ReplaceAllString(base64Str, "")
	decodedData, err := base64.StdEncoding.DecodeString(trimmedBase64Str)
	if err != nil {
		return nil, err
	}
	return decodedData, nil

	img, err := gocv.IMDecode(decodedData, gocv.IMReadUnchanged)
	if err != nil {
		return nil, err
	}
	defer img.Close()

	// 这里假设图片是 RGBA，即含有透明度通道
	if img.Channels() != 4 {
		return decodedData, nil
	}
	bounds := image.Rect(img.Cols()/2, img.Rows()/2, img.Cols()/2, img.Rows()/2)

	for y := 0; y < img.Rows(); y++ {
		for x := 0; x < img.Cols(); x++ {
			// 获取当前像素的 alpha 值
			alpha := img.GetVecbAt(y, x)[3]
			if alpha != 0 {
				if x < bounds.Min.X {
					bounds.Min.X = x
				}
				if x > bounds.Max.X {
					bounds.Max.X = x
				}
				if y < bounds.Min.Y {
					bounds.Min.Y = y
				}
				if y > bounds.Max.Y {
					bounds.Max.Y = y
				}
			}
		}
	}

	// 计算新的宽度和高度
	newWidth := bounds.Max.X - bounds.Min.X + 1
	newHeight := bounds.Max.Y - bounds.Min.Y + 1

	// 如果图片全透明，newWidth或newHeight可能是0
	if newWidth <= 0 || newHeight <= 0 {
		return nil, fmt.Errorf("the image is fully transparent")
	}

	croppedImg := img.Region(bounds)
	defer croppedImg.Close()

	buf, err := gocv.IMEncode(gocv.PNGFileExt, croppedImg)
	if err != nil {
		return nil, err
	}

	return buf.GetBytes(), nil
}
