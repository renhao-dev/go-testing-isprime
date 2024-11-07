package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()

	done := make(chan bool)

	go checkUserInput(os.Stdin, done)

	<-done
	close(done)

	fmt.Println("Thanks!")
}

func checkUserInput(in io.Reader, done chan bool) {
	input := bufio.NewScanner(in)

	for {
		msg, exit := checkNumbers(input)

		if exit {
			done <- true
			return
		}

		fmt.Println(msg)

		prompt()
	}
}

func checkNumbers(input *bufio.Scanner) (string, bool) {
	if !input.Scan() {
		log.Fatal("Input read failed")
		return "", true
	}

	inputText := input.Text()

	if strings.EqualFold(inputText, "q") {
		return "", true
	}

	n, err := strconv.Atoi(inputText)
	if err != nil {
		return "Please, enter a valid number", false
	}

	_, msg := isPrime(n)

	return msg, false
}

func intro() {
	fmt.Println("Is number a prime?")
	fmt.Println("------------------")

	prompt()
}

func prompt() {
	fmt.Print("Enter your number -> ")
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by def", n)
	}
	if n < 0 {
		return false, "Negative numbers are not prime by def"
	}

	for i := 2; i <= n-1; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is divisable by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is prime", n)
}
