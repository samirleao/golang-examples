package imgutil

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
)

// Jpeg jpeg mime type
const Jpeg = "image/jpeg"

// Png png mime type
const Png = "image/png"

// Decode decodes the byte array to an image
func Decode(data []byte) (image.Image, error) {
	contentType := http.DetectContentType(data)
	switch contentType {
	case Jpeg:
		return jpeg.Decode(bytes.NewReader(data))
	case Png:
		return png.Decode(bytes.NewReader(data))
	default:
		return nil, errors.New("Unsuported image type '" + contentType + "'")
	}
}

// Redraw redraws the img into a destination image striping metadata
func Redraw(img image.Image) image.Image {
	newImg := image.NewRGBA(img.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), img, img.Bounds().Min, draw.Over)
	return newImg
}

// Encode encodes the img to a file of the desired contentType type
func Encode(img image.Image, contentType string, file *os.File) error {
	switch contentType {
	case Jpeg:
		opt := new(jpeg.Options)
		opt.Quality = 80
		err := jpeg.Encode(file, img, opt)
		if err != nil {
			return err
		}
	}

	return errors.New("Unsuported contentType '" + contentType + "'")
}
