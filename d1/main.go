package d1

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func Day1() int {
	reader := bufio.NewReader(os.Stdin)

	allCals := make([]int, 0)
	currCals := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			allCals = append(allCals, currCals)
			break
		}

		if line == "\n" {
			allCals = append(allCals, currCals)
			currCals = 0
			continue
		}
		n, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			panic(err)
		}
		currCals += n

	}

	sort.Ints(allCals)
	res := 0
	for i := 0; i < 3; i++ {
		res += allCals[len(allCals)-i-1]
	}

	return res

}
