package d6

import (
	"bufio"
	"os"
	"strings"
)

func Dupes(chars []string) bool {
	vals := make(map[string]bool)

	for _, num := range chars {
		if vals[num] {
			return true
		}
		vals[num] = true
	}
	return false
}

func Day6() int {
	reader := bufio.NewScanner(os.Stdin)
	totalLines := ""

	for reader.Scan() {
		totalLines += reader.Text()
	}
	// Change loop boundaries for Part 2
	for i := 0; i < len(totalLines)-14; i++ {
		group := ""
		for j := i; j < i+14; j++ {
			group = totalLines[i : j+1]
			if len(group) == 14 {
				if !Dupes(strings.Split(group, "")) {
					return j + 1
				}
			}
		}

	}
	return -1
}
