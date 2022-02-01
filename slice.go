package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

var intSlice = make([]int, 3)

// grow append sort and print the slice
func processLoop(input string) {

	i, _ := strconv.Atoi(input)
	intSlice = append(intSlice, int(i))
	sort.Ints(intSlice)
	fmt.Println(intSlice)

}

func main() {

	var input string

	for {

		fmt.Println("Please enter an integer")
		fmt.Scan(&input)

		res, err := regexp.MatchString("^[0-9X]+$", input)

		if err != nil {
			panic(err)
		}

		if res {
			if input == "X" {
				return
			} else {
				processLoop(input)
			}
		} else {
			fmt.Println("That wasn't an integer or an x please try again. Enter x to exit.")
		}

	}
}
