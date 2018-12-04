package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
)

func Process(data []byte) {
	contentType := http.DetectContentType(data)

	switch contentType {
	case "image/jpeg":
		srcImg, err := jpeg.Decode(bytes.NewReader(data))
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		createJpegFile(srcImg, "image.jpg")

	case "image/png":
		srcImg, err := png.Decode(bytes.NewReader(data))
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		createJpegFile(srcImg, "image.jpg")

	default:
		fmt.Println("Unsuported content type " + contentType)
	}
}

func createJpegFile(srcImg image.Image, dir string) {
	newImg := image.NewRGBA(srcImg.Bounds())
	
  // Always replace the background with a white one
  draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
  
  // Draw the source image on the new image
	draw.Draw(newImg, newImg.Bounds(), srcImg, srcImg.Bounds().Min, draw.Over)
	
  file, err := os.Create(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	opt := new(jpeg.Options)
	opt.Quality = 80
	err = jpeg.Encode(file, newImg, opt)
	if err != nil {
		fmt.Println("Error writing jpeg image file.")
		fmt.Println(err)
	}
}
