package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

var nameSlice = make([]name, 0)

func main() {

	var filename string

	fmt.Println("Please enter  filename")
	fmt.Scan(&filename)

	f, fe := os.Open(filename)

	if fe != nil {
		panic(fe)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {

		tokens := strings.Split(sc.Text(), " ")

		s := name{fname: tokens[0], lname: tokens[1]}

		nameSlice = append(nameSlice, s)

	}

	if se := sc.Err(); se != nil {
		fmt.Printf("Scan error: %v", se)
		return
	}

	for _, st := range nameSlice {
		fmt.Println(st.fname, st.lname)
	}

}
