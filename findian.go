package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Please enter a string")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	inputTrimmed := strings.TrimSpace(input)
	inputLower := strings.ToLower(inputTrimmed)

	if err != nil {
		fmt.Println(err)
	}

	if strings.HasPrefix(inputLower, "i") &&
		strings.HasSuffix(inputLower, "n") &&
		strings.Contains(inputLower, "a") {
		println("Found!")
	} else {
		println("Not Found!")
	}
}
