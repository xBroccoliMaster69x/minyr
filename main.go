package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/xBroccoliMaster69x/funtemps/conv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Minyr")
	fmt.Println("Minyr har 2 funksjoner Convert vil generere en .csv fil som inneholder tempereaturverdier i fahrenheit istedenfor celsius")
	fmt.Println("Average gir en gjennomsnittlige temperaturverdier for perioden for kjevik flyplass.")
	fmt.Println("1. Convert")
	fmt.Println("2. Average")

	// Read user input for mode
	fmt.Print("Enter your choice (1 or 2): ")
	modeInput, _ := reader.ReadString('\n')

	// Remove newline character from input
	mode := modeInput[:len(modeInput)-1]
	
	if mode == "1" || mode == "convert" {
		fmt.Println("du har valgt 'Convert'")
		if fileExists("output.txt") {
			fmt.Print("Output fil eksisterer allerede. Regenerer? (j/n): ")
			regenerateInput, _ := reader.ReadString('\n')
			regenerate := regenerateInput[:len(regenerateInput)-1]

			if regenerate == "j" || regenerate == "J" {
				fmt.Println("Regenererer 'output.txt' fil...")
				convert()
			} else {
				fmt.Print("genererer ikke output.txt")
			}

		} else {
			fmt.Println("utdata fil eksisterer ikke genererer output.txt")
			convert()
		}
	
	} else if mode == "2" || mode == "average" {
		fmt.Println("du har valgt 'Average'")
		fmt.Print("velg utdata temperatur type (celsius/fahrenheit): ")
		unitInput, _ := reader.ReadString('\n')
		unit := unitInput[:len(unitInput)-1]
		fmt.Println("du har valgt", unit)
		average(unit)
	} else {
		fmt.Println("Invalid choice. Exiting...")
	}
}

// Helper function to check if a file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func convert() {

	src, err := os.Open("table.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create("output.txt")
	defer dst.Close()
	defer dst.WriteString("Data er basert paa gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Alexander Glasdam Andersen")

	log.Println(src)
	log.Print(dst)

	scanner := bufio.NewScanner(src)
	
	lineCounter := 0
	var fahrenheit float64
	for scanner.Scan() {
		linebuf := scanner.Text()
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
		if lineCounter == 0 {
			dst.WriteString("Navn;Stasjon;Tid(norsk normaltid);Lufttemperatur" + "\n")
			lineCounter ++
			continue
			}
			elementArray := strings.Split(linebuf, ";")
			if len(elementArray) >= 3 {
				celsius := elementArray[3]
				if celsius == ""{
					continue
				}
				celsiusFloat, err := strconv.ParseFloat(celsius, 64)
				if err != nil {
					log.Fatal(err)
				}
				fahrenheit = conv.CelsiusToFahrenheit(celsiusFloat)
				output := "Kjevik;SN39040;" + elementArray[2] + ";" + strconv.FormatFloat(fahrenheit, 'f', 2, 64) + "\n"
				_, err = dst.WriteString(output)
				if err != nil {
					log.Fatal(err)
				}

			}
			
			lineCounter++
		
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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

        avg := tempSum / float64(lineCounter-1)

        if unit == "celsius" {
                log.Printf("Average temperature in Celsius: %.2f", avg)
        } else if unit == "fahrenheit" {
                fahrenheit := conv.CelsiusToFahrenheit(avg) // Use CelsiusToFahrenheit from conv package
                log.Printf("Average temperature in Fahrenheit: %.2f", fahrenheit)
        } else {
                log.Println("Invalid unit. Please specify 'celsius' or 'fahrenheit'")
        }
}
