package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// cow bird snake
// eat move and speak

type IAnimal interface {
	Eat()
	Move()
	Speak()
}

type Animal struct {
	Food, Locomotion, Noise string
}

func (a Animal) Eat() {
	fmt.Println(a.Food)
}

func (a Animal) Move() {
	fmt.Println(a.Locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.Noise)
}

func orchestrate(requestedAnimalName, requestedBehaviour string) {

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	var animal IAnimal

	switch requestedAnimalName {
	case "cow":
		animal = cow
	case "bird":
		animal = bird
	case "snake":
		animal = snake
	default:
		fmt.Println("the requested animal should be one of cow, bird, snake.")

	}

	switch requestedBehaviour {
	case "speak":
		animal.Speak()
	case "move":
		animal.Move()
	case "eat":
		animal.Eat()
	default:
		fmt.Println("the requested beaviour should be one of speak, move, eat.")

	}

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}

		inputTrimmed := strings.TrimSpace(input)
		inputLower := strings.ToLower(inputTrimmed)
		tokens := strings.Split(inputLower, " ")

		if len(tokens) < 2 {
			fmt.Println("There must be two words e.g. cow speak")
		}

		orchestrate(tokens[0], tokens[1])

		f, err := os.Open("test.txt")

		fmt.Println(err)
		fmt.Println(f)
	}
}
