package main

import "fmt"

func main() {
	SRC := "package main\n\nimport \"fmt\"\n\nfunc main() {\n	SRC := \"?\"\n\n	for i := 0; i < len(SRC); i++ {\n		if SRC[i] == 63 {\n			for j := 0; j < len(SRC); j++ {\n				switch SRC[j] {\n				case '\\n':\n					fmt.Printf(\"\\\\n\")\n				case '\"':\n					fmt.Printf(\"\\\\\\\"\")\n				case '\\\\':\n					fmt.Printf(\"\\\\\\\\\")\n				default:\n					fmt.Printf(\"%c\", SRC[j])\n				}\n			}\n		} else {\n			fmt.Printf(\"%c\", SRC[i])\n		}\n	}\n}\n"

	for i := 0; i < len(SRC); i++ {
		if SRC[i] == 63 {
			for j := 0; j < len(SRC); j++ {
				switch SRC[j] {
				case '\n':
					fmt.Printf("\\n")
				case '"':
					fmt.Printf("\\\"")
				case '\\':
					fmt.Printf("\\\\")
				default:
					fmt.Printf("%c", SRC[j])
				}
			}
		} else {
			fmt.Printf("%c", SRC[i])
		}
	}
}
