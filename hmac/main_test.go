package main

import "testing"

func TestSign(t *testing.T) {
	//TODO: write unit test cases for sign()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader
	cases := []struct {
		name 	string
		input string
		expectedOutput string
	}{
		{
			name: "same key",
			input: "a123",
			expectedOutput: 
		},
	}
	for _, c := range cases {

	}

}

func TestVerify(t *testing.T) {
	//TODO: write until test cases for verify()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader
}
