package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	success := false

	reader := bufio.NewReader(os.Stdin)
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have ", 10 - guesses, " guesses left")
		fmt.Println("Make a guess: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)  // delete new line symbol
		guess, err := strconv.Atoi(input) // Atoi is equivalent to ParseInt

		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Your suggest was LOW.")
		} else if guess > target {
			fmt.Println("Your suggest was HIGH.")
		} else {
			success = true
			fmt.Println("Good Job! You guess it!")
			break
		}
	}

	if !success {
		fmt.Println("Sorry you didn't guess my number. It was: ", target)
	}

}
