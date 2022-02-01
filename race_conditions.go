package main

import (
	"fmt"
	//"sync"
	"time"
)

func adder(x *int, routine_name string) {

	*x = 1
	time.Sleep(10000000)
	*x = *x + 1

	fmt.Println(routine_name, *x)
}

func main() {

	var x int

	/*
		The calls to the adder create a race condition becaue the individual instructions
		do not create an atomic integer addition at the machine instruction level from
		perspective of the golang code.

		i.e. x = x + 1 is not an atomic operation.

	*/

	go adder(&x, "routine of 1")
	go adder(&x, "routine of 2")

	time.Sleep(100000000)

	fmt.Println("Complete! : ", x)

}
