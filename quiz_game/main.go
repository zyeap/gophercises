package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// 1. read quiz file in CSV format (default to problems.csv)
// 2. give quiz to user and keep track of right/wrong answers
// 3. keep asking questions regardless of answer
// 4. output total number of Qs user got correct, and total number of Qs

func main() {

	var quiz_file = flag.String("f", "problems.csv", "file name containing quiz game, defaults to problems.csv")

	flag.Parse()

	f, err := os.Open(*quiz_file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	var correct int
	var total int

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		total++
		fmt.Println(rec[0])
		fmt.Print("Answer: ")

		var answer int
		fmt.Scan(&answer)

		solution, err := strconv.Atoi(rec[1])
		if err != nil {
			log.Fatal(err)
		}

		if answer == solution {
			correct++
		}
	}

	fmt.Printf("Num correct: %d\n", correct)
	fmt.Printf("Total questions: %d\n", total)

}
