// Mandelbrot создает PNG-изображение фрактала Мандельброта.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		sampling               = 2
	)

	bufferWidth := width * sampling
	bufferHeight := height * sampling
	buffer := image.NewRGBA(image.Rect(0, 0, bufferWidth, bufferHeight))
	for py := 0; py < bufferHeight; py++ {
		y := float64(py)/float64(bufferHeight)*(ymax-ymin) + ymin
		for px := 0; px < bufferWidth; px++ {
			x := float64(px)/float64(bufferWidth)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Точка (px, py) представляет комплексное значение z.
			buffer.Set(px, py, mandelbrot(z))
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			subpixels := getSubpixels(buffer, sampling, x*sampling, y*sampling)
			c := sample(subpixels)
			img.Set(x, y, c)
		}
	}

	png.Encode(os.Stdout, img) // Примечание: игнорируем ошибки
}

// getSubpixels возвращает суб-пиксели из изображения
func getSubpixels(img image.Image, sampling, x, y int) []*color.Color {
	colors := make([]*color.Color, sampling*sampling)
	index := 0
	for sy := y; sy < y+sampling; sy++ {
		for sx := x; sx < x+sampling; sx++ {
			c := img.At(sx, sy)
			colors[index] = &c
			index++
		}
	}

	return colors
}

// sample возвращает усредненное значение цветов пикселей
func sample(subpixels []*color.Color) color.Color {
	var ar, ag, ab, aa uint64
	for _, c := range subpixels {
		r, g, b, a := (*c).RGBA()
		ar += uint64(r)
		ag += uint64(g)
		ab += uint64(b)
		aa += uint64(a)
	}
	sublen := uint64(len(subpixels))
	r := uint32(ar / sublen)
	g := uint32(ag / sublen)
	b := uint32(ab / sublen)
	a := uint32(aa / sublen)
	result := color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
	}

	return result
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{
				R: contrast*n - 255,
				G: 255 - contrast*n,
				B: 255 - contrast*n,
				A: 255,
			}
		}
	}
	return color.Black
}
