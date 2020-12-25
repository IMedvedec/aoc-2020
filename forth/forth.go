package forth

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	filename = "input/forth_input"
)

var (
	requiredFileds = [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
)

type document map[string]interface{}

func newDocument() document {
	var newDocument document = make(map[string]interface{})
	return newDocument
}

func (doc document) isEmpty() bool {
	return len(doc) == 0
}

func (doc document) isPassport() bool {
	isValid := true

	for _, field := range requiredFileds {
		if _, ok := doc[field]; !ok {
			isValid = false
			break
		}
	}

	return isValid
}

func Run() {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("error while file opening: %v", err)
	}

	var validPasspordCount int

	scanner := bufio.NewScanner(file)
	for document := scanDocument(scanner); !document.isEmpty(); document = scanDocument(scanner) {
		if document.isPassport() {
			validPasspordCount++
		}
	}

	log.Printf("Valid passport count: %d\n", validPasspordCount)
}

func scanDocument(scanner *bufio.Scanner) document {
	document := newDocument()

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		entries := strings.Fields(line)

		for _, entry := range entries {
			pair := strings.Split(entry, ":")
			key, value := pair[0], pair[1]

			document[key] = value
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	return document
}
