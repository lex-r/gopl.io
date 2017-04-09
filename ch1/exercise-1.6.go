// Lissajous генерирует анимированный GIF из случайных
// фигур Лиссажу
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black}

const (
	whiteIndex = 0 // Первый цвет палитры
	blackIndex = 1 // Следующий цвет палитры
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Количество полных колебаний
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // Относительная частота колебаний
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	// Наполняем палитру случайными цветами
	for i := 0; i < 8; i++ {
		palette = append(palette, color.RGBA{
			uint8(rand.Intn(255)), // Яркость красного компонента пикселя
			uint8(rand.Intn(255)), // Яркость зеленого компонента пикселя
			uint8(rand.Intn(255)), // Яркость синего компонента пикселя
			0xff, // Альфа-канал
		})
	}

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(rand.Intn(len(palette)-2) + 1)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += -0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Примечание: игнорируем ошибки
}
