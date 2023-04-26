package go_sikuli

import (
	"bytes"
	"image"
	"image/png"
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
