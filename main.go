package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// docker         | run <image_name> | <cmd> <args> (equals with:)
// go run main.go | run              | <cmd> <args>

type Docker struct {
	name string
	proc []string
}

/* In `C`:
class Wheel {
public:
	std::string material;
	float radius;
};

class Car {
public:
	typedef std::vector<Wheel> WheelCollection; // same as `type DockerImage = Docker` in Golang.
	WheelCollection wheels;
}
*/
type DockerImage = Docker // equivalent to `typedef` in C.

type (
	Mount struct{}
)

// Golang does not provide `class` data type like any other Object-Oriented Language (ie: Java, C++, ...).
// Instead of that, we define associated functions or methods' expressions. Simply, a method is a function with a receiver.
// See: https://go.dev/ref/spec#Method_declarations
func (d *Docker) Name() string {
	return d.name
}

func main() {
	switch os.Args[1] {
	case "run":
		run()

	default:
		panic("unsupported command!")
	}
}

func run() {
	var cmd *exec.Cmd

	// os.Args[2:] := is the command line with all the arguments we expect to run inside the container.
	fmt.Printf("Running cmd %v\n", os.Args[2:])

	// os.Args[2]    := "echo"
	// os.Args[3]... := "-n 'hello The Office'"
	if runtime.GOOS == "windows" {
		fmt.Printf("Detecting OS...: %s\n", runtime.GOOS)

		// windowsCmd := "cmd /c " + os.Args[2]
		// cmd = exec.Command(windowsCmd, os.Args[3:]...)
		windowsCmd := "powershell -NoLogo -NoProfile -NonInteractive " + os.Args[2]
		cmd = exec.Command(windowsCmd, os.Args[3:]...)
		execCmd(cmd)
	}

	cmd = exec.Command(os.Args[2], os.Args[3:]...)
	execCmd(cmd)
}

func execCmd(cmd *exec.Cmd) {
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
