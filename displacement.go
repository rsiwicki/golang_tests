package main

import (
	"fmt"
	"strconv"
)

func GenDisplaceFn(a float64, iv float64, id float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * (t * t)) + (iv * t) + id
	}
}

func main() {

	var input string

	fmt.Println("Please enter accelaration and press return.")
	fmt.Scan(&input)

	a, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Please enter initial velocity and press return.")
	fmt.Scan(&input)

	iv, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Please enter initial displacment and press return.")
	fmt.Scan(&input)

	id, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err)
	}

	fn := GenDisplaceFn(a, iv, id)

	fmt.Println("Please enter the time for displacement.")
	fmt.Scan(&input)

	t, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(fn(t))

}
