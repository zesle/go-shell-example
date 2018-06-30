package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		err = execInput(input)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// ErrNoPath is returned when 'cd' was called without a second argument
var ErrNoPath = errors.New("path required")

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	// Split the input on each space to separate the command and the arguments
	args := strings.Split(input, " ")

	// check for shell build-in command
	switch args[0] {
	case "cd":
		// cd to home with empty path not yet supported
		if len(args) < 2 {
			return ErrNoPath
		}
		err := os.Chdir(args[1])
		if err != nil {
			return err
		}
		// stop further processing
		return nil
	case "exit":
		os.Exit(0)
	}

	// execute programm
	cmd := exec.Command(args[0], args[1:]...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Printf("%s", stdoutStderr)
	return nil
}
