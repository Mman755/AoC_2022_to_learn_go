package d4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Contained(fLower int, fUpper int, sLower int, sUpper int) bool {
	if fLower <= sLower && fUpper >= sUpper {
		return true
	}
	if sLower <= fLower && sUpper >= fUpper {
		return true
	}

	return false
}

func P2Contained(fLower int, fUpper int, sLower int, sUpper int) bool {
	if fLower <= sUpper && sLower <= fUpper {
		return true
	}

	return false
}

// 4-9 and 5-6

func Day4() int {
	reader := bufio.NewScanner(os.Stdin)
	count := 0

	for reader.Scan() {
		line := reader.Text()
		var stringInitial = strings.Split(line, ",")
		var bounds []int

		for i := 0; i < len(stringInitial); i++ {
			currBounds := strings.Split(stringInitial[i], "-")
			num1, _ := strconv.Atoi(currBounds[0])
			num2, _ := strconv.Atoi(currBounds[1])
			bounds = append(bounds, num1)
			bounds = append(bounds, num2)
		}
		if P2Contained(bounds[0], bounds[1], bounds[2], bounds[3]) {
			count += 1
		}

	}

	return count
}
