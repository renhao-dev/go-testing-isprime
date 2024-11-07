package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func Test_isPrime(t *testing.T) {
	tests := []struct {
		name        string
		num         int
		expectedRes bool
		expectedMsg string
	}{
		{"prime", 7, true, "7 is prime"},
		{"not prime", 6, false, "6 is divisable by 2"},
		{"zero", 0, false, "0 is not prime, by def"},
		{"one", 1, false, "1 is not prime, by def"},
		{"negative", -11, false, "Negative numbers are not prime by def"},
	}

	for _, test := range tests {
		res, msg := isPrime(test.num)
		if res != test.expectedRes {
			t.Errorf("%s: Expected result: %t, got: %t", test.name, test.expectedRes, res)
		}
		if msg != test.expectedMsg {
			t.Errorf("%s Expected msg: %s, got: %s", test.name, test.expectedMsg, msg)
		}
	}
}
func Test_intro(t *testing.T) {
	r, w, _ := os.Pipe()

	oldOut := os.Stdout
	os.Stdout = w

	intro()

	_ = w.Close()
	text, _ := io.ReadAll(r)
	output := string(text)

	expectedOutputSubstr := "Enter your number -> "
	if !strings.Contains(output, expectedOutputSubstr) {
		t.Errorf("Incorrect intro text, got '%s'", output)
	}

	os.Stdout = oldOut
}
func Test_prompt(t *testing.T) {
	r, w, _ := os.Pipe()

	oldOut := os.Stdout
	os.Stdout = w

	prompt()

	_ = w.Close()
	text, _ := io.ReadAll(r)
	output := string(text)

	expectedOutput := "Enter your number -> "
	if output != expectedOutput {
		t.Errorf("Invalid output expected: '%s', got '%s'", expectedOutput, output)
	}

	os.Stdout = oldOut
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedMsg    string
		expectedIfDone bool
	}{
		{"prime", "7", "7 is prime", false},
		{"not prime", "6", "6 is divisable by 2", false},
		{"negative", "-11", "Negative numbers are not prime by def", false},
		{"quit", "q", "", true},
		{"not number", "balalalaika.7", "Please, enter a valid number", false},
		{"prime", "7", "7 is prime", false},
	}

	for _, test := range tests {
		input := strings.NewReader(test.input)
		scanner := bufio.NewScanner(input)

		msg, done := checkNumbers(scanner)

		if msg != test.expectedMsg {
			t.Errorf("Expected output msg: %s, got: %s", test.expectedMsg, msg)
		}
		if done != test.expectedIfDone {
			t.Errorf("Expected is done: %t, got: %t", test.expectedIfDone, done)
		}
	}

}

func Test_checkUserInput(t *testing.T) {
	r, w, _ := os.Pipe()

	oldOut := os.Stdout
	os.Stdout = w

	doneChan := make(chan bool)

	var in bytes.Buffer
	in.Write([]byte("7\nq\n"))

	timer := time.NewTimer(time.Millisecond * 100)
	go checkUserInput(&in, doneChan)

	select {
	case <-doneChan:
	case <-timer.C:
		t.Error("Failed to finish")
	}
	close(doneChan)

	_ = w.Close()

	output, _ := io.ReadAll(r)

	if !strings.Contains(string(output), "Enter your number -> ") {
		t.Error("Incorrect ouput")
	}

	os.Stdout = oldOut
}
