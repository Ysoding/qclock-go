package main

import "fmt"

const (
	FontSize = 10
	FontRows = 5
	FontCols = 3
)

var fonts [FontSize]uint32 = [FontSize]uint32{31599, 19812, 14479, 31207, 23524, 29411, 29679, 30866, 31727, 31719}

func main() {

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
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
