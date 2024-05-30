package ascii

import (
	"os"
)

func PrintFlag(inputFile, outputFile *os.File, v []string) {
	
	if Checker(v) {
		for k := 0; k < len(v)-1; k++ {
			outputFile.WriteString("\n")
		}
		return
	}

	for _, word := range v {
		var char []string
	    var chars [][]string
		r := []rune(word)
		if word == "" {

			outputFile.WriteString("\n")

			continue
		}
		for i := 0; i < len(r); i++ {
			char = Printingchar(r[i], inputFile)

			chars = append(chars, char)
		}
		for i := 0; i < 8; i++ {
			for j := 0; j < len(chars); j++ {
				outputFile.WriteString(chars[j][i])
			}

			outputFile.WriteString("\n")

		}
	}
	defer inputFile.Close()
}
