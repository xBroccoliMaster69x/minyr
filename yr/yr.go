package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"path/filepath"
	"os"
	"bufio"
	"strconv"

	"github.com/xBroccoliMaster69x/funtemps/conv"
)

func main() {
	celsiusValues := []float64{6, 0, -11}

	for _, celsius := range celsiusValues {
		fahrenheit := convert(celsius)
		fmt.Printf("%.2f°C is %.2f°F\n", celsius, fahrenheit)
	}

	// Specify the relative file path of kjevik-temp-fahrenheit-20220318-20230318.csv from yr.go
	//inputFilePath := filepath.Join("..","kjevik-temp-celsius-20220318-20230318.csv")
	outputFilePath := filepath.Join("..","kjevik-temp-fahrenheit-20220318-20230318.csv")

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

func average(unit string) float64 {
	inputFilePath := filepath.Join("..","kjevik-temp-celsius-20220318-20230318.csv")
        src, err := os.Open(inputFilePath)
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

        avg := tempSum / float64(lineCounter-1)

        if unit == "celsius" {
                log.Printf("Average temperature in Celsius: %.2f", avg)
        }
	return avg
}
