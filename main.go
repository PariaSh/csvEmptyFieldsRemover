package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
)

func main() {
	file, indices := getParameters()
	updateCSV(file, indices)
}

func getParameters() (string, []int) {
	input := flag.String("input", "", "CSV file whose empty fields need to be removed.")
	indexes := flag.String("index", "", "Indices of fields need to be checked on emptiness; if this is not provided, all fields are going to be checked.")
	flag.Parse()

	if *input == "" {
		log.Fatal("No input CSV file was given.")
	}
	indices := make([]int, 0, 2)
	if *indexes == "" {
		log.Printf("indexes are not specified, all fields are going to be checked.")
	} else {
		for _, i := range strings.Split(*indexes, ",") {
			index, err := strconv.Atoi(strings.TrimSpace(i))
			if err != nil {
				log.Fatalf("Indexes should be the format: \"number[, number]\"; but got: %v", *indexes)
			}
			indices = append(indices, index)
		}
	}
	return *input, indices
}