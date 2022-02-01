package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	m := make(map[string]string)

	var name, address string

	fmt.Println("Please enter your name")
	fmt.Scan(&name)

	m["name"] = name

	fmt.Println("Please enter your address")
	fmt.Scan(&address)

	m["address"] = address

	j, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(j))

}
