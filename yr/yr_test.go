package main

import (
	"log"
	"testing"
	"path/filepath"

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
	outputFilePath := filepath.Join("..", "output.txt")
	want := 27

	amountLines, err := countLines(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if amountLines != want {
		t.Errorf("countLines(%s) = %d, want %d", outputFilePath, amountLines, want)
	}
}

func TestGetLastLine(t *testing.T) {
	outputFilePath := "output.txt"
	want := "Data er basert paa gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av NAVN_PÅ_STUDENTEN"

	lastLine, err := getLastLine(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if lastLine != want {
		t.Errorf("getLastLine(%s) = %s, want %s", outputFilePath, lastLine, want)
	}
}

func TestAverage(t *testing.T) {
	celsiusValues := []float64{6, 0, -11}
	want := 8.56

	average := average(celsiusValues)
	if average != want {
		t.Errorf("average(%v) = %.2f, want %.2f", celsiusValues, average, want)
	}
}
