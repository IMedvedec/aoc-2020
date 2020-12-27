package eight

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	filename   = "input/eight_input"
	commandAcc = "acc"
	commandJmp = "jmp"
	commandNop = "nop"
)

// instruction defines the input instruction.
type instruction struct {
	command string
	offset  int64
}

// newInstruction is a instruction constructor.
func newInstruction(line string) *instruction {
	instuctionParts := strings.Fields(line)

	command := instuctionParts[0]
	offset, err := strconv.ParseInt(instuctionParts[1], 10, 64)
	if err != nil {
		log.Fatalf("number parse error: %v", err)
	}

	return &instruction{
		command: command,
		offset:  offset,
	}
}

// code represents the input list of instructions by row.
type code map[int]*instruction

// newCode is a code constructor.
func newCode() code {
	newCode := make(map[int]*instruction)
	return newCode
}

// findLoop contains the problem solution.
func (c code) findLoop() {
	var accumulator int64

	visitedInstructions := make(map[int]struct{})

	for i := 0; i < len(c); {
		if _, ok := visitedInstructions[i]; ok {
			log.Printf("Loop found, accumulator status: %d\n", accumulator)
			os.Exit(0)
		}

		visitedInstructions[i] = struct{}{}
		instruction := c[i]

		switch instruction.command {
		case commandAcc:
			accumulator += instruction.offset
			i++
		case commandJmp:
			i += int(instruction.offset)
		case commandNop:
			i++
		default:
			log.Fatalf("command type not defined")
		}

	}
}

// Run is the solution starting point.
func Run() {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("error while file opening: %v", err)
	}

	scanner := bufio.NewScanner(file)

	code := getInstructionSet(scanner)

	code.findLoop()
}

// getInstructionSet is the input parser.
func getInstructionSet(scanner *bufio.Scanner) code {
	code := newCode()

	var inputRow int
	for scanner.Scan() {
		line := scanner.Text()

		code[inputRow] = newInstruction(line)

		inputRow++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	return code
}
