package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (a Cow) Eat()   { fmt.Println("grass") }
func (a Cow) Move()  { fmt.Println("walk") }
func (a Cow) Speak() { fmt.Println("moo") }

type Bird struct{}

func (a Bird) Eat()   { fmt.Println("worms") }
func (a Bird) Move()  { fmt.Println("fly") }
func (a Bird) Speak() { fmt.Println("peep") }

type Snake struct{}

func (a Snake) Eat()   { fmt.Println("mice") }
func (a Snake) Move()  { fmt.Println("slither") }
func (a Snake) Speak() { fmt.Println("hsss") }

func ProcessNewAnimal(tokens []string, m map[string]Animal) {

	if len(tokens) < 3 {
		println("command invalid - too few parameters")
		return
	}

	var animal Animal

	animalName := tokens[1]
	animalType := tokens[2]

	switch animalType {
	case "cow":
		animal = Cow{}
	case "bird":
		animal = Bird{}
	case "snake":
		animal = Snake{}
	default:
		fmt.Println("the requested animal should be one of cow, bird, snake.")
		return
	}

	m[animalName] = animal

	fmt.Println("Created it!")

}

func ProcessQuery(tokens []string, m map[string]Animal) {

	if len(tokens) < 3 {
		println("query invalid - too few parameters")
		return
	}

	animalName := tokens[1]
	animalBehaviour := tokens[2]

	var animal Animal

	animal, ok := m[animalName]

	if !ok {
		fmt.Println("animal name not found")
		return
	}

	switch animalBehaviour {
	case "speak":
		animal.Speak()
	case "move":
		animal.Move()
	case "eat":
		animal.Eat()
	default:
		fmt.Println("the requested beaviour should be one of speak, move, eat.")
		return
	}

}

func main() {

	var m map[string]Animal = make(map[string]Animal)

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print(">")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}

		inputTrimmed := strings.TrimSpace(input)
		inputLower := strings.ToLower(inputTrimmed)
		tokens := strings.Split(inputLower, " ")

		switch tokens[0] {
		case "newanimal":
			ProcessNewAnimal(tokens, m)
		case "query":
			ProcessQuery(tokens, m)
		case "default":
			fmt.Println("You must specify a first word of either: query , newanimal")
		}
	}
}
