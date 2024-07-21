package main

import (
	"fmt"
	"time"
)

const (
	FontSize      = 11
	FontRows      = 5
	FontCols      = 3
	DigitsCount   = 8
	DigitsPadding = 2

	DisplayHeight = FontRows
	DisplayWidth  = (FontCols + DigitsPadding) * 8
)

var fonts [FontSize]uint32 = [FontSize]uint32{31599, 19812, 14479, 31207, 23524, 29411, 29679, 30866, 31727, 31719, 1040}

func pixel(ch byte, x int, y int, digits [DigitsCount]int) {
	i := x / (FontCols + DigitsPadding)
	dx := x % (FontCols + DigitsPadding)
	if y < FontRows && dx < FontCols && ((fonts[digits[i]]>>((FontRows-y-1)*FontCols+dx))&1) != 0 {
		fmt.Printf("\033[1;31m%c\033[0m", ch)
	} else {
		fmt.Printf("%c", ch)
	}
}

func main() {
	SRC := "package main\n\nimport (\n	\"fmt\"\n	\"time\"\n)\n\nconst (\n	FontSize      = 11\n	FontRows      = 5\n	FontCols      = 3\n	DigitsCount   = 8\n	DigitsPadding = 2\n\n	DisplayHeight = FontRows\n	DisplayWidth  = (FontCols + DigitsPadding) * 8\n)\n\nvar fonts [FontSize]uint32 = [FontSize]uint32{31599, 19812, 14479, 31207, 23524, 29411, 29679, 30866, 31727, 31719, 1040}\n\nfunc pixel(ch byte, x int, y int, digits [DigitsCount]int) {\n	i := x / (FontCols + DigitsPadding)\n	dx := x % (FontCols + DigitsPadding)\n	if y < FontRows && dx < FontCols && ((fonts[digits[i]]>>((FontRows-y-1)*FontCols+dx))&1) != 0 {\n		fmt.Printf(\"\\033[1;31m%c\\033[0m\", ch)\n	} else {\n		fmt.Printf(\"%c\", ch)\n	}\n}\n\nfunc main() {\n	SRC := \"?\"\n\n	var digits [8]int\n\n	for {\n		currentTime := time.Now()\n		digits[0] = currentTime.Hour() / 10\n		digits[1] = currentTime.Hour() % 10\n		digits[2] = 10\n		digits[3] = currentTime.Minute() / 10\n		digits[4] = currentTime.Minute() % 10\n		digits[5] = 10\n		digits[6] = currentTime.Second() / 10\n		digits[7] = currentTime.Second() % 10\n\n		x := 0\n		y := 0\n		for i := 0; i < len(SRC) && y < FontRows; i++ {\n			switch SRC[i] {\n			case 63:\n				for j := 0; j < len(SRC); j++ {\n					switch SRC[j] {\n					case '\\n':\n						pixel('\\\\', x, y, digits)\n						x += 1\n\n						pixel('n', x, y, digits)\n						x += 1\n					case '\"':\n						pixel('\\\\', x, y, digits)\n						x += 1\n\n						pixel('\"', x, y, digits)\n						x += 1\n					case '\\\\':\n						pixel('\\\\', x, y, digits)\n						x += 1\n\n						pixel('\\\\', x, y, digits)\n						x += 1\n					default:\n						fmt.Printf(\"%c\", SRC[j])\n						x += 1\n					}\n				}\n			case '\\n':\n				y += 1\n				x = 0\n				pixel('\\n', x, y, digits)\n			default:\n				pixel(SRC[i], x, y, digits)\n				x += 1\n			}\n		}\n\n	}\n\n}\n"

	var digits [8]int

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
		for i := 0; i < len(SRC) && y < FontRows; i++ {
			switch SRC[i] {
			case 63:
				for j := 0; j < len(SRC); j++ {
					switch SRC[j] {
					case '\n':
						pixel('\\', x, y, digits)
						x += 1

						pixel('n', x, y, digits)
						x += 1
					case '"':
						pixel('\\', x, y, digits)
						x += 1

						pixel('"', x, y, digits)
						x += 1
					case '\\':
						pixel('\\', x, y, digits)
						x += 1

						pixel('\\', x, y, digits)
						x += 1
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
