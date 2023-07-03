package d7

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
- Track all directories and their corresponding children (Parent <-> Child relationship)
- For each directory have some structure/method that can map the directory/directory name to its size
- From the above find which directories have a size of at most x
*/

var MAX_DIR_SIZE = 100000
var TOTAL_DIR_CAP = 70000000
var MIN_SPACE_REQ = 30000000

func Contains(line []string, lookFor string) bool {
	for i := 0; i < len(line); i++ {
		if line[i] == lookFor {
			return true
		}
	}
	return false
}

func ContainsNumber(line []string) bool {
	regex := regexp.MustCompile("[0-9]")
	for i := 0; i < len(line); i++ {
		if regex.MatchString(line[i]) {
			return true
		}
	}
	return false
}

var parentChild = make(map[string][]string)
var dirSizes = make(map[string]int)

func Day7() int {
	reader := bufio.NewScanner(os.Stdin)
	path := make([]string, 0)
	for reader.Scan() {
		line := reader.Text()
		lineContent := strings.Split(line, " ")

		if Contains(lineContent, "cd") && strings.Join(lineContent, " ") != "$ cd .." {
			path = append(path, lineContent[2])
		}
		if Contains(lineContent, "dir") {
			parentChild[strings.Join(path, "")] = append(parentChild[strings.Join(path, "")], lineContent[1])
		}
		if Contains(lineContent, "..") {
			path = path[:len(path)-1]
		}
		if ContainsNumber(lineContent) {
			val, _ := strconv.Atoi(lineContent[0])
			clonePath := make([]string, 0)
			clonePath = append(clonePath, path...)
			currPath := ""
			for i := 0; i < len(path); i++ {
				currPath += clonePath[i]
				dirSizes[currPath] += val
			}
		}
	}
	constraintSum := 0
	for key := range dirSizes {
		if dirSizes[key] <= MAX_DIR_SIZE {
			constraintSum += dirSizes[key]
		}
	}
	fmt.Println("PART 1", constraintSum)
	fmt.Println("PART 2", P2())

	return P2()

}

func P2() int {
	currDirSize := 0
	currMin := dirSizes["/"]
	freeSpace := TOTAL_DIR_CAP - dirSizes["/"]

	for key := range dirSizes {
		currDirSize = freeSpace + dirSizes[key]
		if currDirSize >= MIN_SPACE_REQ {
			currMin = int(math.Min(float64(dirSizes[key]), float64(currMin)))
		}
	}
	return currMin
}
