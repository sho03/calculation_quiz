package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	hard := flag.Bool("hard", false, "if set, a quiz will be hard mode.")
	flag.Parse()

	var max int
	if *hard {
		max = 10000
	} else {
		max = 100
	}

	left := rand.Intn(max + 1)
	right := rand.Intn(max + 1)

	ops := []string{"+", "-", "*", "/"}
	randomIndex := rand.Intn(len(ops))
	op := ops[randomIndex]

	var answer int
	switch op {
	case "+":
		answer = left + right
	case "-":
		answer = left - right
	case "*":
		answer = left * right
	case "/":
		if left < right {
			temp := left
			left = right
			right = temp
		}
		answer = left / right
	}

	fmt.Println("what is answer ?")
	fmt.Println(left, op, right)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inAnswer, err := strconv.Atoi(sc.Text())

	if err != nil {
		fmt.Println("input number !")
		return
	}

	if answer == inAnswer {
		fmt.Println("You're collect!")
	} else {
		message := fmt.Sprintf("You're wrong! collect answer is %d", answer)
		fmt.Println(message)
	}
}
