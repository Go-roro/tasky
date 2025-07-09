package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/shlex"
	"github.com/yourusername/tasky/cli/cmd"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("📝Tasky - Interactive Task Manager")
	for {
		fmt.Print("tasky > ")
		if !scanner.Scan() {
			break // EOF or Ctrl+D
		}
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}
		if line == "exit" || line == "quit" {
			fmt.Println("👋 Goodbye!")
			break
		}

		args, err := shlex.Split(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, "❌ Invalid input:", err)
			continue
		}
		if err := cmd.ExecuteCommand(args); err != nil {
			fmt.Fprintln(os.Stderr, "❌", err)
		}
	}
}
