package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	// A command that prints slowly, so we can see streaming happen.
	cmd := exec.Command("sh", "-c", "for i in 1 2 3 4 5; do echo line $i; sleep 1; done")

	// Ask for a pipe to its output BEFORE starting it.
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("pipe failed:", err)
		return
	}

	// Start runs the program but does NOT wait for it.
	if err := cmd.Start(); err != nil {
		fmt.Println("start failed:", err)
		return
	}

	// Read whatever has arrived, over and over, until the program closes the pipe.
	buf := make([]byte, 1024)
	for {
		n, err := stdout.Read(buf)
		if n > 0 {
			fmt.Printf("[chunk of %d bytes] %s", n, string(buf[:n]))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read failed:", err)
			break
		}
	}

	// Wait collects the exit code, now that output is finished.
	if err := cmd.Wait(); err != nil {
		fmt.Println("exited with error:", err)
	}
	fmt.Println("--- done, exit code:", cmd.ProcessState.ExitCode(), "---")
}
