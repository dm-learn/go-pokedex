package main

import (
	"math"
	"testing"
)

func TestCalculateCatchProbability(t *testing.T) {
	tests := map[string]struct {
		input int
		expected float64
	} {
		"zero_exp": {input: 0, expected: 0.8},
		"min_probability": {input: 10000, expected: 0.1},
		"exp_800": {input: 800, expected: 0.4},
		"exp_431": {input: 431, expected: 0.5845},
		"exp_1": {input: 1, expected: 0.7995},
	}

	for name, c := range tests {
		actual := calculateCatchProbability(c.input)
		if math.Abs(actual - c.expected) > 0.00000001 {
			t.Errorf("Test failed: %s. Expected: %v. Actual %v", name, c.expected, actual)
		}
	}
}

func TestCalculateCatchThreshold(t *testing.T) {
	tests := map[string]struct {
		input int
		expected int
	} {
		"1": {input: 1, expected: 0},
		"100": {input: 100, expected: 75},
		"700": {input: 700, expected: 315},
		"1600": {input: 1600, expected: 160},
		"1400": {input: 1400, expected: 140},
	}

	for name, c := range tests {
		actual := calculateCatchThreshold(c.input)
		if actual != c.expected {
			t.Errorf("Test failed: %s. Expected: %v. Actual %v", name, c.expected, actual)
		}
	}
}