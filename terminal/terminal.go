package terminal

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dwiw96/komerce-coding-test/psbb"
	sortcharacter "github.com/dwiw96/komerce-coding-test/sort-character"
)

const menu = `1. Sort Character
2. PSBB
3. Menu
4. Exit
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
			fmt.Println("\n-- psbb --")
			fmt.Print("<< enter number of families: ")
			families := readFamilies()
			if families == 0 {
				fmt.Println("families must be greater than 0")
				continue
			}
			fmt.Print("<< enter number of member in the family: ")
			members := readMembers(families)

			psbb.MinimumBus(families, members)
			fmt.Println()
		case "3":
			fmt.Print(menu)
		case "4":
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

func readFamilies() int {
	var families int
	_, err := fmt.Scan(&families)
	if err != nil {
		fmt.Println(">> error:", err)
	}

	return families
}

func readMembers(families int) []int {
	members := make([]int, families, families*2)

	for i := 0; i < families; i++ {
		_, err := fmt.Scan(&members[i])
		if err != nil {
			fmt.Println(">> error:", err)
			return nil
		}
		if members[i] > 4 {
			fmt.Println(">> maximum member of familiy is 4")
			return nil
		}
	}

	if families != len(members) {
		fmt.Println("Input must be equal with count of family")
	}

	return members
}
