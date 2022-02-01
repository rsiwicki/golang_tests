package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(pos int, target []int) {
	l := target[pos]
	//r := target[pos+1]
	target[pos] = target[pos+1]
	target[pos+1] = l
}

func bubbleSort(target []int) {
	for i := 0; i < len(target); i++ {
		for j := 0; j < (len(target) - i - 1); j++ {
			if target[j] > target[j+1] {
				swap(j, target)
			}
		}
	}
}

func main() {

	//var intSlice = []int{2, 3, 7, 7, 2, 4, 4, 7, 7, 5, 6}

	fmt.Println("Please enter series of space delimited integers and press return e.g. 1 4 6 7 3")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	stringSlice := strings.Split(strings.TrimSpace(input), " ")[:]

	var intSlice = []int{}

	for _, i := range stringSlice {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, j)
	}

	// so there is no need to pass-by-reference per say because the
	// slice is passed and it still refers to the underlying array by pointer
	bubbleSort(intSlice)

	fmt.Println(intSlice)

}
