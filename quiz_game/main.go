package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// 1. read quiz file in CSV format (default to problems.csv)
// 2. give quiz to user and keep track of right/wrong answers
// 3. keep asking questions regardless of answer
// 4. output total number of Qs user got correct, and total number of Qs

func main() {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", rec[0])
	}

}
