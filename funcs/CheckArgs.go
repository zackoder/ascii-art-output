package ascii

import (
	"fmt"
	"os"
	"strings"
)

func CheckArgs(args []string) ([]string, []string) {
	if len(args) > 4 || len(args) <= 1 || (len(args) == 2 && strings.Contains(args[1], "--output")) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}

	if len(args) == 3 && strings.Contains(args[1], "--output") {
		args = append(args, "standard")
	}
	if len(args) == 2 && !strings.Contains(args[1], "--output") {
		args = append(args, "standard")
	} else if len(args) == 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(1)
	}

	input := args[len(args)-2]
	r := []rune(input)
	for i := 0; i < len(r); i++ {
		if r[i] < 32 ||
			r[i] > 127 {
			fmt.Println("you need to choose a character from the ascii table.")
			os.Exit(2)
		}
	}
	words := strings.Split(input, "\\n")
	if len(args) == 4 {
		v := strings.Split(args[1], "=")

		if len(v[0]) != 8 {
			fmt.Println("You chose the incorrect flag.")
			os.Exit(3)
		} else if len(v) > 1 && !strings.HasSuffix(v[1], ".txt") {
			fmt.Println("The file name is incorrect")
			os.Exit(4)
		}
	}
	return words, args
}
