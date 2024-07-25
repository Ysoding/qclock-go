package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// TODO: use like std_c_lexer generate format

	inputPath := "./qclock.go"
	outputPath := "./qclock_blob.go"

	inputData, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("error reading input file %s: %v\n", inputPath, err)
	}

	output, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("could not create file %s: %v\n", inputPath, err)
	}
	defer output.Close()

	for _, ch := range inputData {
		if ch == '?' {
			for _, innerChar := range inputData {
				switch innerChar {
				case '\n':
					fmt.Fprintf(output, "\\n\"\n\"")
				case '\\':
					fmt.Fprint(output, "\\\\")
				case '"':
					fmt.Fprint(output, "\\\"")
				default:
					fmt.Fprintf(output, "%c", innerChar)
				}
			}
		} else {
			fmt.Fprintf(output, "%c", ch)
		}
	}
}
