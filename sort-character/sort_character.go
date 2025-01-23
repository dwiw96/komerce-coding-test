package sortcharacter

import (
	"fmt"
	"strings"
)

func SortCharacter(s string) {
	vowelMap := map[string][2]int{
		"a": {},
		"i": {},
		"u": {},
		"e": {},
		"o": {},
	}

	vowelCount := make(map[string]int)
	vowelOrder := []string{}
	consonantCount := make(map[string]int)
	consonantOrder := []string{}

	// temp := [][]int32{}
	// var vowelCount int32 = 0

	s = removeSpace(s)
	s = strings.ToLower(s)

	for _, char := range s {
		if !((string(char) >= "a" && string(char) <= "z") || (string(char) >= "A" && string(char) <= "Z")) {
			continue
		}
		if _, isOk := vowelMap[string(char)]; isOk {
			if _, isOk := vowelCount[string(char)]; isOk {
				vowelCount[string(char)]++
				continue
			}
			vowelOrder = append(vowelOrder, string(char))
			vowelCount[string(char)]++
			// if slices.Index(temp[], string(char)) != -1 {}
			// vowelCount++
			// temp = append(temp, []int32{char, vowelCount})
			continue
		}
		if _, isOk := consonantCount[string(char)]; isOk {
			consonantCount[string(char)]++
			continue
		}
		consonantOrder = append(consonantOrder, string(char))
		consonantCount[string(char)]++
	}

	vowel := placingCharacters(vowelOrder, vowelCount)
	consonant := placingCharacters(consonantOrder, consonantCount)

	fmt.Println(">> Vowel:", vowel)
	fmt.Println(">> Consonant:", consonant)
}

func removeSpace(s string) (result string) {
	for _, char := range s {
		if string(char) == " " {
			continue
		}
		result += string(char)
	}
	return
}

func placingCharacters(order []string, count map[string]int) (res string) {
	for _, val := range order {
		c := count[val]
		for i := 0; i < c; i++ {
			res += val
			count[val]--
		}
	}
	return res
}
