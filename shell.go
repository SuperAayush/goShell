package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/chzyer/readline"
)

func main() {
	r1, err := readline.New(">> ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Readline Error:", err)
		return
	}

	defer r1.Close()

	for {
		input, err := r1.Readline()

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}

}

var errorPath = errors.New("path required")

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errorPath
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
