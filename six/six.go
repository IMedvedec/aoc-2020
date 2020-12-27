package six

import (
	"bufio"
	"log"
	"os"
)

const (
	filename = "input/six_input"
)

type questions map[rune]struct{}

func newQuestions() questions {
	var newQuestions questions = make(map[rune]struct{})
	return newQuestions
}

func (q questions) isEmpty() bool {
	return len(q) == 0
}

func (q questions) count() int {
	return len(q)
}

func Run() {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("error while file opening: %v", err)
	}

	var questionSum int
	scanner := bufio.NewScanner(file)

	for questions := analyseGroup(scanner); !questions.isEmpty(); questions = analyseGroup(scanner) {
		questionSum += questions.count()
	}

	log.Printf("Question sum: %d\n", questionSum)
}

func analyseGroup(scanner *bufio.Scanner) questions {
	questions := newQuestions()

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		for _, r := range line {
			if _, ok := questions[r]; !ok {
				questions[r] = struct{}{}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	return questions
}
