package main

import (
	"fmt"
	"os"
	"strings"
)

type Input struct {
	input string
	file  *os.File
}

const (
	SELECT = "SELECT"
	INSERT = "INSERT"
	UPDATE = "UPDATE"
)

func NewInput() *Input {
	return &Input{}
}

func (i *Input) ReadInput() {
	fmt.Print("> ")
	_, err := fmt.Scanln(&i.input)
	if err != nil {
		fmt.Println("Error reading input: ", err)
	}

	prompt, err := i.handleInput()

	if err != nil {
		if err.Error() == "exit" {
			fmt.Println(prompt)
			os.Exit(0)
		}

		fmt.Println("Error handling input: ", err)
		return
	}

	fmt.Println(prompt)
}

func (i *Input) handleInput() (string, error) {
	i.input = strings.ToUpper(i.input)

	switch i.input {
	case ".EXIT":
		return "Goodbye!", fmt.Errorf("exit")
	case SELECT:
		return "This is where we would do a SELECT.", nil
	case INSERT:
		return "This is where we would do an INSERT.", nil
	case UPDATE:
		return "This is where we would do an UPDATE.", nil
	default:
		return "Unrecognized command: " + i.input, fmt.Errorf("Unrecognized command: %s", i.input)
	}
}
