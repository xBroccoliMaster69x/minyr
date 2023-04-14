package main

import (
	"log"
	"testing"
	"path/filepath"
	"os"
	"math"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 6, want: 42.8},
		{input: 0, want: 32.0},
		{input: -11, want: 12.2},
	}

	for _, tc := range tests {
		got := convert(tc.input)
		if got != tc.want {
			t.Errorf("convert(%.2f) = %.2f, want %.2f", tc.input, got, tc.want)
		}
	}
}

func TestCountLines(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	outputFilePath := filepath.Join(wd, "../kjevik-temp-fahrenheit-20220318-20230318.csv") // Use os.Getwd() to get the current working directory and construct the relative file path
	want := 16756

	amountLines, err := countLines(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if amountLines != want {
		t.Errorf("countLines(%s) = %d, want %d", outputFilePath, amountLines, want)
	}
}

func TestGetLastLine(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	outputFilePath := filepath.Join(wd, "../kjevik-temp-fahrenheit-20220318-20230318.csv") // Use os.Getwd() to get the current working directory and construct the relative file path
	want := "Data er basert paa gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Alexander Glasdam Andersen"

	lastLine, err := getLastLine(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if lastLine != want {
		t.Errorf("getLastLine(%s) = %s, want %s", outputFilePath, lastLine, want)
	}
}
func TestAverage(t *testing.T) {
	unit := "celsius"
	expected := 8.56
	tolerance := 0.005 // set a tolerance value for comparison

	// Call the average function and get the calculated average
	avg := average(unit)

	// Compare the calculated average with the expected value within the tolerance
	if math.Abs(avg-expected) > tolerance {
		t.Errorf("Average temperature is incorrect. Expected: %.2f, got: %.2f", expected, avg)
	}
}

//func TestAverage(t *testing.T) {
//	unit := "celsius"
//	expected := 8.56
//
//	// Call the average function and get the calculated average
//	avg := average(unit)
//
//	// Compare the calculated average with the expected value
//	if avg != expected {
//		t.Errorf("Average temperature is incorrect. Expected: %.2f, got: %.2f", expected, avg)
//	}
//}
