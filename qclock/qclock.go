package main

import (
	"fmt"
	"time"
)

var digits [8]int
var fonts [11]uint32 = [11]uint32{31599, 19812, 14479, 31207, 23524, 29411, 29679, 30866, 31727, 31719, 1040}

func pixel(ch byte, x int, y int, digits [8]int) {
	i := x / 5
	dx := x % 5
	if i < 8 && y < 5 && dx < 3 && ((fonts[digits[i]]>>((5-y-1)*3+dx))&1) != 0 {
		fmt.Printf("\033[1;31m%c\033[0m", ch)
	} else {
		fmt.Printf("%c", ch)
	}
}

func main() {
	SRC := "?"

	for {
		currentTime := time.Now()
		digits[0] = currentTime.Hour() / 10
		digits[1] = currentTime.Hour() % 10
		digits[2] = 10
		digits[3] = currentTime.Minute() / 10
		digits[4] = currentTime.Minute() % 10
		digits[5] = 10
		digits[6] = currentTime.Second() / 10
		digits[7] = currentTime.Second() % 10

		x := 0
		y := 0
		for i := 0; i < len(SRC) && y < 3; i++ {
			switch SRC[i] {
			case 63:
				for j := 0; j < len(SRC); j++ {
					switch SRC[j] {
					case '\n':
					case '\\':
					case '"':
						pixel('\\', x, y, digits)
						x += 1
						fallthrough
					default:
						pixel(SRC[j], x, y, digits)
						x += 1
					}
				}
			case '\n':
				y += 1
				x = 0
				pixel('\n', x, y, digits)
			default:
				pixel(SRC[i], x, y, digits)
				x += 1
			}

		}
		fmt.Printf("\033[%dA\033[%dD", y, x)
		time.Sleep(1 * time.Second)
	}
}
