package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"path/filepath"

	"github.com/xBroccoliMaster69x/funtemps/conv"
)

func main() {
	celsiusValues := []float64{6, 0, -11}

	for _, celsius := range celsiusValues {
		fahrenheit := convert(celsius)
		fmt.Printf("%.2f°C is %.2f°F\n", celsius, fahrenheit)
	}

	// Specify the relative file path of output.txt from yr.go
	outputFilePath := filepath.Join("..","output.txt")

	amountLines, err := countLines(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of lines in %s: %d\n", outputFilePath, amountLines)

	lastLine, err := getLastLine(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Last line in %s: %s\n", outputFilePath, lastLine)

	average := average(celsiusValues)
	fmt.Printf("Average of Celsius values: %.2f°C\n", average)
}

func convert(celsius float64) float64 {
	return conv.CelsiusToFahrenheit(celsius)
}

func countLines(filename string) (int, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	return len(strings.Split(string(content), "\n")), nil
}

func getLastLine(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(content), "\n")
	if len(lines) > 0 {
		return lines[len(lines)-1], nil
	}
	return "", fmt.Errorf("file %s is empty", filename)
}

//func average(values []float64) float64 {
//	sum := 0.0
//	for _, v := range values {
//		sum += v
//	}
//	return sum / float64(len(values))
//}
