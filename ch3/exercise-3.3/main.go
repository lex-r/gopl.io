// Surface вычисляет SVG-представление трехмерного графика функции.
package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320            // Размер канвы в пикселях
	cells         = 100                 // Количество ячеек сетки
	xyrange       = 30.0                // Диапазон осей (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // Пикселей в единице x и y
	zscale        = height * 0.4        // Пикселей в единице z
	angle         = math.Pi / 6         // Углы осей x, y (=30 градусов)
	minz          = -0.3                // Минимальная высота для вычисления цвета ячейки
	maxz          = 0.3                 // Максимальная высота для вычисления цвета ячейки
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30), cos(30)
var (
	mincolor = color.RGBA{255, 0, 0, 255} // Цвет ячейки минимальной высоты
	maxcolor = color.RGBA{0, 0, 255, 255} // Цвет ячейки максимальной высоты
)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(pos(i+1, j))
			bx, by := corner(pos(i, j))
			cx, cy := corner(pos(i, j+1))
			dx, dy := corner(pos(i+1, j+1))

			_, _, z := pos(i+1, j)
			c := zcolor(minz, maxz, z, mincolor, maxcolor)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style=\"fill:%s;\"/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, colortostr(c))
		}
	}

	fmt.Println("</svg>")
}

func pos(i, j int) (float64, float64, float64) {
	// Ищем угловую точку (x,y) ячейки (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхности z
	z := f(x, y)

	return x, y, z
}

func corner(x, y, z float64) (float64, float64) {
	// Изометрически проецируем (x,y,z) на двумерную канву SVG (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func zcolor(minz, maxz, z float64, minColor, maxColor color.RGBA) color.RGBA {
	if z < minz {
		return minColor
	}

	if z > maxz {
		return maxColor
	}

	length := maxz - minz
	zpos := z - minz
	zposperc := zpos / length
	c := color.RGBA{}
	rdiff := zposperc * (float64(maxColor.R) - float64(minColor.R))
	c.R = uint8(float64(minColor.R) + rdiff)
	gdiff := zposperc * (float64(maxColor.G) - float64(minColor.G))
	c.G = uint8(float64(minColor.G) + gdiff)
	bdiff := zposperc * (float64(maxColor.B) - float64(minColor.B))
	c.B = uint8(float64(minColor.B) + bdiff)

	return c
}

func colortostr(c color.RGBA) string {
	r := fmt.Sprintf("%x", c.R)
	if len(r) == 1 {
		r = "0" + r
	}

	g := fmt.Sprintf("%x", c.G)
	if len(g) == 1 {
		g = "0" + g
	}

	b := fmt.Sprintf("%x", c.B)
	if len(b) == 1 {
		b = "0" + b
	}

	return fmt.Sprintf("#%s%s%s", r, g, b)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Расстояние от (0,0)
	if r == 0 {
		return 0
	}

	return math.Sin(r) / r
}
