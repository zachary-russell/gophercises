package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"flag"
)

func main() {
	input := flag.String("file", "foo.csv", "CSV to Input")
	flag.Parse()
	csvfile, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvfile))
	right := 0
	wrong := 0
	for {
		line, error := reader.Read()
		var text string
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		fmt.Println("What is " + line[0] + "?")
		fmt.Scanln(&text)
		if text == line[1] {
			right = right + 1;
		} else {
			fmt.Println("false")
			wrong = wrong + 1;
		}
	}
	fmt.Println("Correct Answers:")
	fmt.Println(right)
	fmt.Println("Wrong Answers:")
	fmt.Println(wrong)
}

