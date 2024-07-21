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
)

var fonts [FontSize]uint32 = [FontSize]uint32{31599, 19812, 14479, 31207, 23524, 29411, 29679, 30866, 31727, 31719, 1040}

func printFontChar() {
	// 111
	// 001
	// 010
	// 010
	// 010

	// => 111 001 010 010 010
	// => 29330
	for i := 0; i < FontSize; i++ {
		for y := 0; y < FontRows; y++ {
			for x := 0; x < FontCols; x++ {
				v := fonts[i] >> ((FontRows-y-1)*FontCols + x)
				if v&1 != 0 {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func main() {
	currentTime := time.Now()

	var digits [8]int

	digits[0] = currentTime.Hour() / 10
	digits[1] = currentTime.Hour() % 10
	digits[2] = 10
	digits[3] = currentTime.Minute() / 10
	digits[4] = currentTime.Minute() % 10
	digits[5] = 10
	digits[6] = currentTime.Second() / 10
	digits[7] = currentTime.Second() % 10
	fmt.Printf("%02d:%02d:%02d\n", currentTime.Hour(), currentTime.Minute(), currentTime.Second())

	h := FontRows
	w := (FontCols + DigitsPadding) * 8
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			i := x / (FontCols + DigitsPadding)
			dx := x % (FontCols + DigitsPadding)
			if (dx < FontCols) && (((fonts[digits[i]] >> ((FontRows-y-1)*FontCols + dx)) & 1) != 0) {
				// fmt.Printf("\x1b[33m#\x1b[0m")
				// fmt.Printf("\x1b[31m#\x1b[0m")
				fmt.Printf("\033[1;31m█\033[0m")
			} else {
				fmt.Printf("█")
			}
		}
		fmt.Println()
	}
}
