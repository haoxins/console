package console

import "runtime"
import "strings"
import "os/exec"
import "fmt"
import "os"

var platform string
var sttyFlag string

func init() {
	platform = runtime.GOOS

	switch platform {
	case "darwin":
		sttyFlag = "-f"
	case "linux":
		sttyFlag = "-F"
	}
}

func InterceptChar() string {
	// disable input buffer
	e := exec.Command("stty", sttyFlag, "/dev/tty", "cbreak", "min", "1").Run()
	exitIfError(e)
	// do not display input characters on the screen
	e = exec.Command("stty", sttyFlag, "/dev/tty", "-echo").Run()
	exitIfError(e)

	defer exec.Command("stty", sttyFlag, "/dev/tty", "echo").Run()

	b := make([]byte, 1)
	os.Stdin.Read(b)

	return string(b[:])
}

func InterceptLine() string {
	// disable input buffer
	e := exec.Command("stty", sttyFlag, "/dev/tty", "cbreak", "min", "1").Run()
	exitIfError(e)
	// do not display input characters on the screen
	e = exec.Command("stty", sttyFlag, "/dev/tty", "-echo").Run()
	exitIfError(e)

	defer exec.Command("stty", sttyFlag, "/dev/tty", "echo").Run()

	buffer := new([]byte)

	for {
		b := make([]byte, 1)
		os.Stdin.Read(b)

		if int(b[0]) == 10 {
			break
		}

		*buffer = append(*buffer, b[0])
	}

	return string((*buffer)[:])
}

func ReadChar() string {
	// disable input buffer
	e := exec.Command("stty", sttyFlag, "/dev/tty", "cbreak", "min", "1").Run()
	exitIfError(e)

	b := make([]byte, 1)
	os.Stdin.Read(b)

	return string(b[:])
}

func ReadLine() string {
	var s string
	fmt.Scanf("%s", &s)
	return strings.TrimSpace(s)
}

func exitIfError(e error) {
	if e != nil {
		fmt.Printf("error: %s", e.Error())
		os.Exit(1)
	}
}
