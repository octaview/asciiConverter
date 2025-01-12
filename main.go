package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"

	xdraw "golang.org/x/image/draw"
)

const (
	asciiChars  = "@%#*+=-:."
	outputWidth = 100
)

func main() {
	imgPath := "image.jpg" // Укажите путь к изображению

	file, err := os.Open(imgPath)
	if err != nil {
		fmt.Println("Ошибка открытия изображения:", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Ошибка декодирования изображения:", err)
		return
	}

	scaledImg := resizeImage(img, outputWidth)

	asciiArt := imageToASCII(scaledImg)
	fmt.Println(asciiArt)
}

func resizeImage(img image.Image, targetWidth int) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	aspectRatio := float64(height) / float64(width)
	targetHeight := int(float64(targetWidth) * aspectRatio * 0.6)

	scaledRect := image.Rect(0, 0, targetWidth, targetHeight)
	scaledImg := image.NewRGBA(scaledRect)
	xdraw.ApproxBiLinear.Scale(scaledImg, scaledRect, img, bounds, draw.Over, nil)

	return scaledImg
}

func imageToASCII(img image.Image) string {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	var result strings.Builder

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			gray := pixelToGray(img.At(x, y))
			char := grayToChar(gray)
			result.WriteString(char)
		}
		result.WriteString("\n")
	}

	return result.String()
}

func pixelToGray(c color.Color) uint8 {
	r, g, b, _ := c.RGBA()
	return uint8(0.299*float64(r/256) + 0.587*float64(g/256) + 0.114*float64(b/256))
}

func grayToChar(gray uint8) string {
	if gray > 240 {
		return " "
	}
	scale := float64(len(asciiChars)-1) / 255.0
	return string(asciiChars[int(float64(gray)*scale)])
}
