package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/yourusername/tasky/cli/cmd"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("📝tasky - Interactive Task Manager")
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

		args := strings.Fields(line) // split input into tokens
		if err := cmd.ExecuteCommand(args); err != nil {
			fmt.Fprintln(os.Stderr, "❌", err)
		}
	}
}
