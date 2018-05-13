package main

import (
	"bufio"
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

		err = execLine(input)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func execLine(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	if len(args) < 1 {
		args = append(args, "")
	}
	cmd := exec.Command(args[0], args[1:]...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Printf("%s", stdoutStderr)
	return nil
}
