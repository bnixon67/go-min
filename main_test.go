package main

import (
	"testing"
)

func TestMin(t *testing.T) {
	testCases := []struct {
		name           string
		inputs         []int
		expectedOutput int
	}{
		{
			name:           "Single input",
			inputs:         []int{5},
			expectedOutput: 5,
		},
		{
			name:           "Multiple inputs",
			inputs:         []int{7, 2, 4, 1},
			expectedOutput: 1,
		},
		{
			name:           "One negative",
			inputs:         []int{7, -2, 4, 1},
			expectedOutput: -2,
		},
		{
			name:           "Alll negative",
			inputs:         []int{-7, -2, -4, -1},
			expectedOutput: -7,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := min(tc.inputs[0], tc.inputs[1:]...)
			if output != tc.expectedOutput {
				t.Errorf("unexpected output for inputs %v: got %d, want %d", tc.inputs, output, tc.expectedOutput)
			}
		})
	}
}
