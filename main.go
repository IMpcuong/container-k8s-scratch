package main

import (
	"fmt"
	"os"
	"os/exec"
)

// docker         | run <image_name> | <cmd> <args> (equals with:)
// go run main.go | run              | <cmd> <args>

func main() {
	switch os.Args[1] {
	case "run":
		run()

	default:
		panic("unsupported command!")
	}
}

func run() {
	// os.Args[2:] := is the command line with all the arguments we expect to run inside the container.
	fmt.Printf("Running cmd %v\n", os.Args[2:])

	// os.Args[2]    := "echo"
	// os.Args[3]... := "hello The Office"
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
