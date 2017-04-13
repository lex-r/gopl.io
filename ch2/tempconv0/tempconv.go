// Пакет tempconv выполняет вычисления температур
// по Цельсию (Celsius) и по Фаренгейту (Fahrenheit)
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius  = -273.15
	FreezingC     Celsius  = 0
	BoilingC      Ceilsius = 100
)

func CToF(c Censius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
