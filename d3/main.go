package d3

import (
	"bufio"
	"os"
	"strings"
)

func Contains(arr1 []string, arr2 []string, char string) bool {
	var inFirst bool
	var inSecond bool
	for _, v := range arr1 {
		if v == char {
			inFirst = true
		}
	}
	for _, x := range arr2 {
		if x == char {
			inSecond = true
		}
	}
	inBoth := inFirst && inSecond

	return inBoth
}

func alphaIndex(alphabet [26]string, char string) int {
	for i := 0; i < len(alphabet); i++ {
		if char == alphabet[i] {
			return i
		}
	}
	return -1
}

func Day3() int {
	reader := bufio.NewScanner(os.Stdin)

	var alphabet [26]string

	for i := 0; i < 26; i++ {
		alphabet[i] = string(rune('a' + i))
	}

	sum := 0

	for reader.Scan() {
		line1 := reader.Text()
		reader.Scan()
		line2 := reader.Text()
		reader.Scan()
		line3 := reader.Text()

		var first []string = make([]string, len(line1))
		var second []string = make([]string, len(line2))
		var third []string = make([]string, len(line3))

		for i := 0; i < len(line1); i++ {
			first[i] = string(line1[i])
		}
		for i := 0; i < len(line2); i++ {
			second[i] = string(line2[i])
		}
		for i := 0; i < len(line3); i++ {
			third[i] = string(line3[i])
		}

		for _, char := range first {
			if Contains(second, third, char) {
				if char == strings.ToUpper(char) {
					sum += len(alphabet) + alphaIndex(alphabet, strings.ToLower(char)) + 1
					break
				}
				if char == strings.ToLower(char) {
					sum += alphaIndex(alphabet, char) + 1
					break
				}
			}
		}

	}

	return sum

}
