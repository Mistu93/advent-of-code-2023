package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	log.Printf("Strating reading youg Elf bug mess")
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatalf("File error %s", err.Error())
	}

	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += getNumber(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("The sercret code was: %d", sum)

}

func getNumber(line string) int {
	r, _ := regexp.Compile(`\D*(\d?).*(\d)\D*`)
	matches := r.FindAllStringSubmatch(line, -1)

	var strNumber string
	if len(matches[0][1]) > 0 {
		strNumber = fmt.Sprintf("%s%s", matches[0][1], matches[0][2])
	} else {
		strNumber = fmt.Sprintf("%s%s", matches[0][2], matches[0][2])
	}

	result, err := strconv.Atoi(strNumber)

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	return result
}
