package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

// fmt.Scan also reads input from standard input device
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
	timeLimit := flag.Int("timer", 10, "")
	shuffle := flag.Bool("shuffle", false, "")
	flag.Parse()

	var questions map[string]int

	if *fileName != "problems.csv" {
		questions = readFile(*fileName)
	} else {
		questions = readFile()
	}

	correctAnswers := 0

	if *shuffle {
		rand.Seed(time.Now().UnixNano())
		// todo: fix shuffling function
		rand.Shuffle(len(questions), func(i int, j int) {})
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

QuestionLoop:
	for k, v := range questions {
		select {
		case <-timer.C:
			break QuestionLoop
		default:
		}

		fmt.Printf("%s: ", k)

		answer := make(chan int)

		go func() {
			ans := input()
			answer <- ans
		}()

		select {
		case <-timer.C:
			break QuestionLoop
		case ans := <-answer:
			if ans == v {
				correctAnswers++
			}
		}
	}

	fmt.Printf("\nYou got %d out of %d questions\n", correctAnswers, len(questions))
}
