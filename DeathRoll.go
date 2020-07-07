package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func interupt(x int) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press Enter to Roll for -> (1 - ", x, ")")
	reader.ReadString('\n')
}
func display(x int) {
	if x == 1 {
		fmt.Println("You Roll", x, "\nYou Lose...")
	} else {
		fmt.Println("You Roll:", x)
		interupt(x)
	}
}
func roll(x int) int {
	for x > 1 {

		rand.Seed(time.Now().UnixNano())
		x = rand.Intn(x) + 1
		display(x)
	}
	return x
}

func main() {

	if len(os.Args) > 1 {
		arg := os.Args[1]
		/** converting the arg1 variable into an int using Atoi method (normal int) */
		startRoll, err := strconv.Atoi(arg)
		if err == nil {
			fmt.Println("Starting Range: 1 -", startRoll)
			roll(startRoll)
		} else {
			fmt.Println("Error Parsing Input")
			// Wrong input.
			os.Exit(1)
		}
	} else {
		fmt.Println("No arguments passed, Rolling from 100")
		roll(100)
	}
}
