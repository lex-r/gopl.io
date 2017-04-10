package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // Первый цвет палитры
	blackIndex = 1 // Следующий цвет палитры
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	var (
		cycles  = 5     // Количество полных колебаний
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)

	if err := r.ParseForm(); err == nil {
		cycles = getRequestValue(r, "cycles", cycles)
		size = getRequestValue(r, "size", size)
		nframes = getRequestValue(r, "nframes", nframes)
		delay = getRequestValue(r, "delay", delay)
	}

	lissajous(w, cycles, size, nframes, delay, res)
}

func lissajous(out io.Writer, cycles, size, nframes, delay int, res float64) {
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // Относительная частота колебаний
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += -0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Примечание: игнорируем ошибки
}

// Возвращает значение параметра из запроса, если такой есть, иначе значение по умолчанию
func getRequestValue(r *http.Request, paramName string, defaultValue int) int {
	value, ok := r.Form[paramName]
	if !ok {
		return defaultValue
	}

	if len(value) != 1 {
		return defaultValue
	}

	v, err := strconv.Atoi(value[0])
	if err != nil {
		return defaultValue
	}

	return v
}
