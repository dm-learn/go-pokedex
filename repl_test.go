package main

import (
    "testing"
)


func TestCleanInput(t *testing.T) {
    tests := map[string]struct {
        input string
	    expected []string
    } {
        "helloWorld": {input: " hello world ", expected: []string{"hello", "world"}},
        "noSpace": {input: "HelloWorld", expected: []string{"helloworld"}},
        "allSpaces": {input: " ", expected: []string{""}},
	    "pokemon": {input: "Charmander Bulbasaur PIKACHU", expected: []string{"charmander", "bulbasaur", "pikachu"}},
    }

    for name, c := range tests {
        actual := cleanInput(c.input)
        for i, word := range actual {
            if word != c.expected[i] {
                t.Errorf("Test failed: %v. Expected: %v. Actual: %v", name, c.expected[i], word)
            }
        }
    }
}
