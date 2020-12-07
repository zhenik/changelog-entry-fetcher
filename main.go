package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Argument
	fptr := flag.String("path", "", "File path to read from")
	flag.Parse()
	// Validation
	if !strings.HasSuffix(*fptr, ".md") {
		log.Fatalf("failed opening file: %s", errors.New("input file must have markdown format"))
	}

	// Try to open file
	file, err := os.Open(*fptr)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	// Parse
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	// h1 entry starts with #
	// h2 entry starts with ##

	var h1amount int = 0
	var h2amount int = 0

	for scanner.Scan() {
		line := scanner.Text()

		if h2amount < 2 {
			if strings.HasPrefix(line, "## ") {
				h2amount += 1
			}

			if h2amount < 2 && h1amount == 1 {
				txtlines = append(txtlines, line)
			}

			if strings.HasPrefix(line, "# ") {
				h1amount += 1
			}
		}
	}

	file.Close()

	for _, line := range txtlines {
		fmt.Println(line)
	}

}
