package main

import (
	"log"
	"testing"
	"path/filepath"
	"os"
	"bufio"
	"strings"
	"strconv"
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

	outputFilePath := filepath.Join(wd, "../output.txt") // Use os.Getwd() to get the current working directory and construct the relative file path
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
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	outputFilePath := filepath.Join(wd, "../output.txt") // Use os.Getwd() to get the current working directory and construct the relative file path
	want := "Data er basert paa gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Alexander Glasdam Andersen"

	lastLine, err := getLastLine(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if lastLine != want {
		t.Errorf("getLastLine(%s) = %s, want %s", outputFilePath, lastLine, want)
	}
}
func average(unit string) {
        src, err := os.Open("table.csv")
        if err != nil {
                log.Fatal(err)
        }
        defer src.Close()

        scanner := bufio.NewScanner(src)

        lineCounter := 0
        tempSum := 0.0
        for scanner.Scan() {
                linebuf := scanner.Text()
                if scanner.Err() != nil {
                        log.Fatal(scanner.Err())
                }
                if lineCounter == 0 {
                        lineCounter ++
                        continue
                }
                //if linebuf[len(linebuf)-1] != ';' {
			elementArray := strings.Split(linebuf, ";")
                        if len(elementArray) > 3 {
                                celsius := elementArray[3]
                                if celsius == "" {
                                        continue
                                }
                                celsiusFloat, err := strconv.ParseFloat(celsius, 64)
                                if err != nil {
                                        log.Fatal(err)
                                }
                                tempSum += celsiusFloat
                                lineCounter++

                        }

                
        }
        if err := scanner.Err(); err != nil {
                log.Fatal(err)
        }

        avg := tempSum / float64(lineCounter)

        if unit == "celsius" {
                log.Printf("Average temperature in Celsius: %.2f", avg)
        }
}
