package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestConvertToFahrenheit(t *testing.T) {
	tests := []struct {
		input       string
		want        string
		lineCounter int
	}{
		{input: "6", want: "42.8", lineCounter: 27},
		{input: "0", want: "32.0", lineCounter: 27},
		{input: "-11", want: "12.2", lineCounter: 27},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := ConvertToFahrenheit(tc.input)
			if got != tc.want {
				t.Errorf("ConvertToFahrenheit(%s) = %s, want %s", tc.input, got, tc.want)
			}

			if lineCounter != tc.lineCounter {
				t.Errorf("lineCounter = %d, want %d", lineCounter, tc.lineCounter)
			}

			// Open the output.txt file for reading
			file, err := os.Open("output.txt")
			if err != nil {
				t.Fatalf("Failed to open output.txt: %v", err)
			}
			defer file.Close()

			// Read the file line by line
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, "Data er basert paa gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Alexander Glasdam Andersen") {
					// Found the expected string in the final line
					return
				}
			}

			// If the expected string was not found, fail the test
			t.Errorf("Expected string not found in the final line of output.txt")
		})
	}

	// Check the average value
	avg := Average()
	expectedAvg := 8.56
	if avg != expectedAvg {
		t.Errorf("Average() = %f, want %f", avg, expectedAvg)
	}
}
