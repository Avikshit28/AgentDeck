package main

import (
	"fmt"
	
	"os/exec"

	"github.com/creack/pty"
)

func main() {
	cmd := exec.Command("sh", "-c", "if [ -t 1 ]; then echo 'I am talking to a terminal'; else echo 'I am talking to a pipe'; fi")

	// Start the command attached to a pseudo-terminal.
	// ptmx is our end of it: we read what the program prints,
	// and we can write keystrokes into it.
	ptmx, err := pty.Start(cmd)
	if err != nil {
		fmt.Println("pty start failed:", err)
		return
	}
	defer ptmx.Close()

	// Exactly the same reading loop as before.
	buf := make([]byte, 1024)
	for {
		n, err := ptmx.Read(buf)
		if n > 0 {
			fmt.Printf("[chunk of %d bytes] %s", n, string(buf[:n]))
		}
		if err != nil {
			break
		}
	}

	cmd.Wait()
	fmt.Println("\n--- done ---")
}