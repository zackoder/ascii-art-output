package ascii

import (
	"fmt"
	"os"
)

func PrintN(inputfile *os.File, v []string) {
	
	if Checker(v) {
		for k := 0; k < len(v)-1; k++ {
			fmt.Println()
		}
		return
	}

	for _, word := range v {
		var char []string
		var chars [][]string
		r := []rune(word)
		if word == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < len(r); i++ {
			char = Printingchar(r[i], inputfile)
			chars = append(chars, char)
		}

		for l := 0; l < 8; l++ {
			for j := 0; j < len(chars); j++ {
				fmt.Printf("%s", chars[j][l])
			}
			fmt.Println()
		}
	}
	defer inputfile.Close()
}
