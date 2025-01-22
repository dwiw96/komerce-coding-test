package terminal

import (
	"bufio"
	"fmt"
	"os"

	sortcharacter "github.com/dwiw96/komerce-coding-test/sort-character"
)

const menu = `1. Sort Character
2. PSBB
3. Exit
`

func RunApp() {
	var input string
	fmt.Print(menu)
	for {
		fmt.Print("\nin: ")

		fmt.Scan(&input)

		switch input {
		case "1":
			fmt.Println("\n-- sort character --")
			fmt.Print("<< enter the input: ")
			input = readInput()
			sortcharacter.SortCharacter(input)
		case "2":
			fmt.Println(">> still not available")
		case "3":
			return
		}
	}
}

func readInput() string {
	var strInput string

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		strInput = scanner.Text()
	}

	return strInput
}
