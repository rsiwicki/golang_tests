package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// The question is not explicitly asking for a merge sort
func partialSort(wg *sync.WaitGroup, sortMe []int) {

	fmt.Println("Sorting: ", sortMe)

	sort.Ints(sortMe)

	fmt.Println("Sorted: ", sortMe)

	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	//var sortMe = []int{1, 6, 7, 21, 47, 2, 822, 7, 7, 2, 20, 20}

	fmt.Println("Please enter series of space delimited integers and press return e.g. 1 4 6 7 3")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	stringSlice := strings.Split(strings.TrimSpace(input), " ")[:]

	var sortMe = []int{}

	for _, i := range stringSlice {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		sortMe = append(sortMe, j)
	}

	l := len(sortMe)

	if l < 4 {
		fmt.Println("The question does ask for four goroutines thus length is not less than 4")
	}

	partitionSize := l / 4
	fmt.Println("length is: ", l)
	fmt.Println("partition size is: ", partitionSize)

	// half array then half - this is a bit weird becuase the question explicitly
	// ask for 4 susbsets thus the minimum array size can only be 4
	partition1 := sortMe[0:(partitionSize * 1)]
	partition2 := sortMe[(partitionSize * 1):(partitionSize * 2)]
	partition3 := sortMe[(partitionSize * 2):(partitionSize * 3)]
	partition4 := sortMe[(partitionSize * 3):l]

	wg.Add(4)
	go partialSort(&wg, partition1)
	go partialSort(&wg, partition2)
	go partialSort(&wg, partition3)
	go partialSort(&wg, partition4)
	wg.Wait()

	res := append(append(append(partition1, partition2...), partition3...), partition4...)

	fmt.Println("Partial sort result: ", res)

	sort.Ints(res)

	fmt.Println("Sort the partial sort (makes previous redundant but hey!)", res)

}
