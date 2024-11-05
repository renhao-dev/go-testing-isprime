package main

import "testing"

func Test_isPrime(t *testing.T) {
	tests := []struct {
		name        string
		num         int
		expectedRes bool
		expectedMsg string
	}{
		{"prime", 7, true, "7 is prime"},
		{"prime", 6, true, "6 is prime"},
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
