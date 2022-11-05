package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func input() int {
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return 0
	}

	// remove the delimiter from the string
	input = strings.TrimSuffix(input, "\n")

	integer, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Please enter a valid number!")
		return -1
	}

	return integer
}

func readFile(fileName ...string) map[string]int {
	var f *os.File
	var err error

	if len(fileName) == 0 {
		f, err = os.Open("problems.csv")
	} else {
		f, err = os.Open(fileName[0])
	}

	if err != nil {
		log.Fatalf("Invalid file")
	}

	reader := csv.NewReader(f)
	reader.Comment = '#'
	question := make(map[string]int)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		question[record[0]], _ = strconv.Atoi(record[1])
	}

	return question
}

func main() {
	fileName := flag.String("filename", "problems.csv", "")
	flag.Parse()

	var questions map[string]int

	if *fileName != "problems.csv" {
		questions = readFile(*fileName)
	} else {
		questions = readFile()
	}

	correctAnswers := 0

	for k, v := range questions {
		fmt.Printf("%s: ", k)
		answer := input()

		if answer == v {
			correctAnswers++
		}
	}

	fmt.Printf("\nYou got %d out of %d questions\n", correctAnswers, len(questions))
}
