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
	answer := rand.Intn(100) + 1

	fmt.Println("I have picked a number between 1 and 100.")

	reader := bufio.NewReader(os.Stdin)

	success := false

	for remaining := 10; remaining > 0; remaining-- {
		fmt.Println("Guesses remaining:", remaining)

		fmt.Print("Take a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input) // Atoi returns int
		if err != nil {
			log.Fatal(err)
		}

		if guess < answer {
			fmt.Println("Sorry, that is too low.")
			continue
		} else if guess > answer {
			fmt.Println("Sorry, that is too high")
			continue
		} else {
			fmt.Println("Well done, you guessed", answer, "which is correct!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("You ran out of guesses, the answer was ", answer)
	}
}
