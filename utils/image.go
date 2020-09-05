package utils

import (
	"bytes"
	_ "golang.org/x/image/webp"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
)

func ImageToJPG(rawData []byte) ([]byte, error) {
	r := bytes.NewReader(rawData)
	img, _, err := image.Decode(r)
	w := &bytes.Buffer{}
	err = jpeg.Encode(w, img, &jpeg.Options{Quality: 90})
	return w.Bytes(), err
}
