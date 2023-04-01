package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"github.com/xBroccoliMaster69x/funtemps/conv"
)	

func main() {
	src, err := os.Open("table.csv")
	//src, err := os.Open("/home/janisg/minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
        	log.Fatal(err)
	}
	defer src.Close()
        log.Println(src)
        
	
	scanner := bufio.NewScanner(src)
	linebuf := ""
	for scanner.Scan() {
		linebuf += scanner.Text()
		if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
		if linebuf[len(linebuf)-1] == '\n' {
			log.Println(linebuf)
			elementArray := strings.Split(linebuf,";")
			if len(elementArray) > 3 {
				celsius := elementArray[3]
				fahr := conv.CelsiusToFahrenheit(celsius)
				log.Println(elementArray[3])
			}
			linebuf = ""
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
